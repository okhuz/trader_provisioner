package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeletionHandler(w http.ResponseWriter, r *http.Request) {
	var tm Trader
	dr := DeletionRequest{}
	_ = json.NewDecoder(r.Body).Decode(&dr)
	dbFindResult := db.instance.Where("user_id = ? AND trader_id = ?", dr.UserId, dr.TraderId).First(&tm)
	if dbFindResult.Error != nil {
		log.Println(dbFindResult.Error.Error())
		http.Error(w, dbFindResult.Error.Error(), http.StatusBadRequest)
	} else {
		dbDeleteResult := db.instance.Delete(&tm)
		if dbDeleteResult.Error != nil {
			msg := fmt.Sprintf("could not delete trader for user %s, with ID of %s", tm.UserId, tm.TraderId)
			log.Println(msg)
			http.Error(w, msg, http.StatusBadRequest)
		} else {
			log.Println("deleting deployment")
			deletePolicy := metav1.DeletePropagationForeground
			deploymentsInterface, configMapInterface, serviceInterface, ingressInterface := createClientSets()

			if err := deploymentsInterface.Delete(context.TODO(), config.TraderPrefix+tm.TraderId, metav1.DeleteOptions{
				PropagationPolicy: &deletePolicy,
			}); err != nil {
				log.Println(err)
			}
			if err := configMapInterface.Delete(context.TODO(), config.TraderPrefix+tm.TraderId, metav1.DeleteOptions{
				PropagationPolicy: &deletePolicy,
			}); err != nil {
				log.Println(err)
			}
			if err := serviceInterface.Delete(context.TODO(), config.TraderPrefix+tm.TraderId, metav1.DeleteOptions{
				PropagationPolicy: &deletePolicy,
			}); err != nil {
				log.Println(err)
			}
			if err := ingressInterface.Delete(context.TODO(), config.TraderPrefix+tm.TraderId, metav1.DeleteOptions{
				PropagationPolicy: &deletePolicy,
			}); err != nil {
				log.Println(err)
			}

			log.Println("deleted deployment")
		}
	}

}
