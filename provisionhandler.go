package main

import (
	"context"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"log"
	"net/http"
)

func ProvisionHandler(w http.ResponseWriter, r *http.Request) {
	pr := ProvisionRequest{}
	_ = json.NewDecoder(r.Body).Decode(&pr)
	deploymentId := uuid.NewV4().String()
	resourceIdentifier := traderPrefix + deploymentId
	deploymentsClient, configMapClient, serviceClient := createClientSets()

	deployment := createDeployment(resourceIdentifier)
	configMap := createConfigMap(resourceIdentifier, pr.Config)
	service := createService(resourceIdentifier)

	trader := Trader{Name: pr.Name, TraderId: deploymentId, TradingModel: pr.TradingModel}
	dbResult := db.instance.Create(&trader)
	if dbResult.Error != nil {
		log.Printf(dbResult.Error.Error())
		http.Error(w, dbResult.Error.Error(), http.StatusBadRequest)
		return
	} else {
		log.Println("Record has been created")
		log.Println("Creating deployment...")
		// TODO: Better error handling here maybe array of errors if any revert the changes
		deploymentResult, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
		configMapResult, _ := configMapClient.Create(context.TODO(), configMap, metav1.CreateOptions{})
		serviceResult, _ := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
		if err != nil {
			log.Println(err)
		}
		log.Printf("Created deployment %q.\n", deploymentResult.GetObjectMeta().GetName())
		log.Printf("Created config map %q.\n", configMapResult.GetObjectMeta().GetName())
		log.Printf("Created config map %q.\n", serviceResult.GetObjectMeta().GetName())
		response, _ := json.Marshal(ProvisionResponse{ Id: deploymentId })
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func createDeployment(resourceIdentifier string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: resourceIdentifier,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"trader": resourceIdentifier,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"trader": resourceIdentifier,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "trader-config",
									MountPath: "/freqtrade/user_data/config.json",
									SubPath:   "config.json",
								},
								{
									Name: "strategies-pvc",
									MountPath: "/freqtrade/user_data/strategies",
								},
								{
									Name: "notebooks-pvc",
									MountPath: "/freqtrade/user_data/notebooks",
								},
							},
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
					Volumes: []apiv1.Volume{
						{
							Name: "trader-config",
							VolumeSource: apiv1.VolumeSource{
								ConfigMap: &apiv1.ConfigMapVolumeSource{
									LocalObjectReference: apiv1.LocalObjectReference{
										Name: resourceIdentifier,
									},
									DefaultMode: pointer.Int32Ptr(0777),
								},
							},
						},
						{
							Name: "notebooks-pvc",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: "notebooks-pvc",
								},
							},
						},
						{
							Name: "strategies-pvc",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: "strategies-pvc",
								},
							},
						},
					},
				},
			},
		},
	}
}

func createConfigMap(resourceIdentifier string, config string) *apiv1.ConfigMap {
	return &apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: resourceIdentifier,
		},
		Data: map[string]string{
			"config.json": config,
		},
	}
}

func createService(resourceIdentifier string) *apiv1.Service {
	return &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: resourceIdentifier,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"trader": resourceIdentifier,
			},
			Ports: []apiv1.ServicePort {
				{
					Protocol: apiv1.ProtocolTCP,
					Port: 80,
				},
			},
		},

	}
}
