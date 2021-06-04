# Services
List of services that trader provisioner accepts.
## Provision a Trader 
```POST``` /provision

Request
```json
{
	"user_id": "user_id",
	"configuration": "{\r\n    \"max_open_trades\": 10,\r\n    \"stake_currency\": \"USDT\",\r\n    \"stake_amount\": \"unlimited\",\r\n    \"tradable_balance_ratio\": 0.99,\r\n    \"fiat_display_currency\": \"USD\",\r\n    \"timeframe\": \"1h\",\r\n    \"dry_run\": true,\r\n    \"cancel_open_orders_on_exit\": true,\r\n    \"unfilledtimeout\": {\r\n        \"buy\": 10,\r\n        \"sell\": 30\r\n    },\r\n    \"bid_strategy\": {\r\n        \"ask_last_balance\": 0.0,\r\n        \"use_order_book\": false,\r\n        \"order_book_top\": 1,\r\n        \"check_depth_of_market\": {\r\n            \"enabled\": false,\r\n            \"bids_to_ask_delta\": 1\r\n        }\r\n    },\r\n    \"ask_strategy\": {\r\n        \"use_order_book\": false,\r\n        \"order_book_min\": 1,\r\n        \"order_book_max\": 1,\r\n        \"use_sell_signal\": true,\r\n        \"sell_profit_only\": false,\r\n        \"ignore_roi_if_buy_signal\": false\r\n    },\r\n    \"exchange\": {\r\n        \"name\": \"binance\",\r\n        \"key\": \"MUEI9XWUexwPHR6zWu3o3VLZ0sY2eQ021BEXGBz2t5f77OX5XlwbONpSP2trQeKi\",\r\n        \"secret\": \"h2b0o7EMwf3qHFHEtTFPXkqHCF9BkCGKSil1qx4oUHL1alwVWejilCSTUkABWnRd\",\r\n        \"ccxt_config\": {\"enableRateLimit\": true},\r\n        \"ccxt_async_config\": {\r\n            \"enableRateLimit\": true,\r\n            \"rateLimit\": 200\r\n        },\r\n        \"pair_whitelist\": [\r\n            \"BTC\/USDT\",\r\n            \"1INCH\/USDT\",\r\n            \"AAVE\/USDT\",\r\n            \"ACM\/USDT\",\r\n            \"ADA\/USDT\",\r\n            \"AION\/USDT\",\r\n            \"AKRO\/USDT\",\r\n            \"ALGO\/USDT\",\r\n            \"ALICE\/USDT\",\r\n            \"ALPHA\/USDT\",\r\n            \"ANKR\/USDT\",\r\n            \"ANT\/USDT\",\r\n            \"ARDR\/USDT\",\r\n            \"ARPA\/USDT\",\r\n            \"ASR\/USDT\",\r\n            \"ATM\/USDT\",\r\n            \"ATOM\/USDT\",\r\n            \"AUDIO\/USDT\",\r\n            \"AUD\/USDT\",\r\n            \"AUTO\/USDT\",\r\n            \"AVA\/USDT\",\r\n            \"AVAX\/USDT\",\r\n            \"AXS\/USDT\",\r\n            \"BADGER\/USDT\",\r\n            \"BAL\/USDT\",\r\n            \"BAND\/USDT\",\r\n            \"BAT\/USDT\",\r\n            \"BCH\/USDT\",\r\n            \"BEAM\/USDT\",\r\n            \"BEL\/USDT\",\r\n            \"BLZ\/USDT\",\r\n            \"BNT\/USDT\",\r\n            \"BTCST\/USDT\",\r\n            \"BTT\/USDT\",\r\n            \"BUSD\/USDT\",\r\n            \"BZRX\/USDT\",\r\n            \"CAKE\/USDT\",\r\n            \"CELO\/USDT\",\r\n            \"CELR\/USDT\",\r\n            \"CFX\/USDT\",\r\n            \"CHR\/USDT\",\r\n            \"CHZ\/USDT\",\r\n            \"CKB\/USDT\",\r\n            \"COCOS\/USDT\",\r\n            \"COMP\/USDT\",\r\n            \"COS\/USDT\",\r\n            \"COTI\/USDT\",\r\n            \"CRV\/USDT\",\r\n            \"CTK\/USDT\",\r\n            \"CTSI\/USDT\",\r\n            \"CTXC\/USDT\",\r\n            \"CVC\/USDT\",\r\n            \"DASH\/USDT\",\r\n            \"DATA\/USDT\",\r\n            \"DCR\/USDT\",\r\n            \"DEGO\/USDT\",\r\n            \"DENT\/USDT\",\r\n            \"DGB\/USDT\",\r\n            \"DIA\/USDT\",\r\n            \"DNT\/USDT\",\r\n            \"DOCK\/USDT\",\r\n            \"DODO\/USDT\",\r\n            \"DOGE\/USDT\",\r\n            \"DOT\/USDT\",\r\n            \"DREP\/USDT\",\r\n            \"DUSK\/USDT\",\r\n            \"EGLD\/USDT\",\r\n            \"ENJ\/USDT\",\r\n            \"EOS\/USDT\",\r\n            \"EPS\/USDT\",\r\n            \"ETC\/USDT\",\r\n            \"EUR\/USDT\",\r\n            \"FET\/USDT\",\r\n            \"FIL\/USDT\",\r\n            \"FIO\/USDT\",\r\n            \"FIRO\/USDT\",\r\n            \"FIS\/USDT\",\r\n            \"FLM\/USDT\",\r\n            \"FTM\/USDT\",\r\n            \"FTT\/USDT\",\r\n            \"FUN\/USDT\",\r\n            \"GBP\/USDT\",\r\n            \"GRT\/USDT\",\r\n            \"GTO\/USDT\",\r\n            \"GXS\/USDT\",\r\n            \"HARD\/USDT\",\r\n            \"HBAR\/USDT\",\r\n            \"HIVE\/USDT\",\r\n            \"HNT\/USDT\",\r\n            \"HOT\/USDT\",\r\n            \"ICX\/USDT\",\r\n            \"INJ\/USDT\",\r\n            \"IOST\/USDT\",\r\n            \"IOTA\/USDT\",\r\n            \"IOTX\/USDT\",\r\n            \"IRIS\/USDT\",\r\n            \"JST\/USDT\",\r\n            \"JUV\/USDT\",\r\n            \"KAVA\/USDT\",\r\n            \"KEY\/USDT\",\r\n            \"KMD\/USDT\",\r\n            \"KNC\/USDT\",\r\n            \"KSM\/USDT\",\r\n            \"LINA\/USDT\",\r\n            \"LINK\/USDT\",\r\n            \"LIT\/USDT\",\r\n            \"LRC\/USDT\",\r\n            \"LSK\/USDT\",\r\n            \"LTC\/USDT\",\r\n            \"LTO\/USDT\",\r\n            \"LUNA\/USDT\",\r\n            \"MANA\/USDT\",\r\n            \"MATIC\/USDT\",\r\n            \"MBL\/USDT\",\r\n            \"MDT\/USDT\",\r\n            \"MFT\/USDT\",\r\n            \"MITH\/USDT\",\r\n            \"MKR\/USDT\",\r\n            \"MTL\/USDT\",\r\n            \"NANO\/USDT\",\r\n            \"NBS\/USDT\",\r\n            \"NEAR\/USDT\",\r\n            \"NEO\/USDT\",\r\n            \"NMR\/USDT\",\r\n            \"NULS\/USDT\",\r\n            \"OCEAN\/USDT\",\r\n            \"OGN\/USDT\",\r\n            \"OG\/USDT\",\r\n            \"OMG\/USDT\",\r\n            \"OM\/USDT\",\r\n            \"ONE\/USDT\",\r\n            \"ONG\/USDT\",\r\n            \"ONT\/USDT\",\r\n            \"ORN\/USDT\",\r\n            \"OXT\/USDT\",\r\n            \"PAXG\/USDT\",\r\n            \"PAX\/USDT\",\r\n            \"PERL\/USDT\",\r\n            \"PERP\/USDT\",\r\n            \"PNT\/USDT\",\r\n            \"POND\/USDT\",\r\n            \"PSG\/USDT\",\r\n            \"PUNDIX\/USDT\",\r\n            \"QTUM\/USDT\",\r\n            \"RAMP\/USDT\",\r\n            \"REEF\/USDT\",\r\n            \"REP\/USDT\",\r\n            \"RIF\/USDT\",\r\n            \"RLC\/USDT\",\r\n            \"ROSE\/USDT\",\r\n            \"RSR\/USDT\",\r\n            \"RUNE\/USDT\",\r\n            \"RVN\/USDT\",\r\n            \"SAND\/USDT\",\r\n            \"SC\/USDT\",\r\n            \"SFP\/USDT\",\r\n            \"SKL\/USDT\",\r\n            \"SNX\/USDT\",\r\n            \"SOL\/USDT\",\r\n            \"SRM\/USDT\",\r\n            \"STMX\/USDT\",\r\n            \"STORJ\/USDT\",\r\n            \"STPT\/USDT\",\r\n            \"STRAX\/USDT\",\r\n            \"STX\/USDT\",\r\n            \"SUN\/USDT\",\r\n            \"SUPER\/USDT\",\r\n            \"SUSD\/USDT\",\r\n            \"SUSHI\/USDT\",\r\n            \"SXP\/USDT\",\r\n            \"TCT\/USDT\",\r\n            \"TFUEL\/USDT\",\r\n            \"THETA\/USDT\",\r\n            \"TKO\/USDT\",\r\n            \"TOMO\/USDT\",\r\n            \"TRB\/USDT\",\r\n            \"TROY\/USDT\",\r\n            \"TRU\/USDT\",\r\n            \"TRX\/USDT\",\r\n            \"TUSD\/USDT\",\r\n            \"TWT\/USDT\",\r\n            \"UMA\/USDT\",\r\n            \"UNFI\/USDT\",\r\n            \"UNI\/USDT\",\r\n            \"USDC\/USDT\",\r\n            \"UTK\/USDT\",\r\n            \"VET\/USDT\",\r\n            \"VITE\/USDT\",\r\n            \"VTHO\/USDT\",\r\n            \"WAN\/USDT\",\r\n            \"WAVES\/USDT\",\r\n            \"WING\/USDT\",\r\n            \"WIN\/USDT\",\r\n            \"WNXM\/USDT\",\r\n            \"WRX\/USDT\",\r\n            \"WTC\/USDT\",\r\n            \"XEM\/USDT\",\r\n            \"XLM\/USDT\",\r\n            \"XMR\/USDT\",\r\n            \"XRP\/USDT\",\r\n            \"XTZ\/USDT\",\r\n            \"XVS\/USDT\",\r\n            \"YFII\/USDT\",\r\n            \"YFI\/USDT\",\r\n            \"ZEC\/USDT\",\r\n            \"ZEN\/USDT\",\r\n            \"ZIL\/USDT\",\r\n            \"ZRX\/USDT\",\r\n            \"BNB\/USDT\",\r\n            \"ETH\/USDT\",\r\n            \"REN\/USDT\",\r\n            \"NKN\/USDT\",\r\n            \"BTS\/USDT\"\r\n        ],\r\n        \"pair_blacklist\": [\r\n            \"BNB\/BTC\"\r\n        ]\r\n    },\r\n    \"pairlists\": [\r\n        {\"method\": \"StaticPairList\"}\r\n    ],\r\n    \"edge\": {\r\n        \"enabled\": false,\r\n        \"process_throttle_secs\": 3600,\r\n        \"calculate_since_number_of_days\": 7,\r\n        \"allowed_risk\": 0.01,\r\n        \"stoploss_range_min\": -0.0001,\r\n        \"stoploss_range_max\": -0.1,\r\n        \"stoploss_range_step\": -0.0001,\r\n        \"minimum_winrate\": 0.60,\r\n        \"minimum_expectancy\": 0.20,\r\n        \"min_trade_number\": 10,\r\n        \"max_trade_duration_minute\": 1440,\r\n        \"remove_pumps\": false\r\n    },\r\n    \"telegram\": {\r\n        \"enabled\": false,\r\n        \"token\": \"1777355043:AAFV61iE0FgXEkyZEbOdEYfsjldGI3cFehE\",\r\n        \"chat_id\": \"1574109528\"\r\n    },\r\n    \"api_server\": {\r\n        \"enabled\": true,\r\n        \"listen_ip_address\": \"0.0.0.0\",\r\n        \"listen_port\": 8080,\r\n        \"verbosity\": \"error\",\r\n        \"jwt_secret_key\": \"somethingrandom\",\r\n        \"CORS_origins\": [],\r\n        \"username\": \"freqtrader\",\r\n        \"password\": \"SuperSecurePassword\"\r\n    },\r\n    \"bot_name\": \"turantrade_bot\",\r\n    \"initial_state\": \"running\",\r\n    \"forcebuy_enable\": false,\r\n    \"internals\": {\r\n        \"process_throttle_secs\": 5\r\n    },\r\n    \"strategy\": \"PortfolioStrategy\"\r\n}"
}
```
Response
```json
{
  "id": "Trader/Pod ID"
}
```

## Update Trader Configuration

```PUT``` /update
```json
{
	"user_id": "user_id",
	"trader_id": "trader_id",
	"configuration": "{\r\n    \"max_open_trades\": 3,\r\n    \"stake_currency\": \"USDT\",\r\n    \"stake_amount\": \"unlimited\",\r\n    \"tradable_balance_ratio\": 0.99,\r\n    \"fiat_display_currency\": \"USD\",\r\n    \"timeframe\": \"1h\",\r\n    \"dry_run\": true,\r\n    \"cancel_open_orders_on_exit\": false,\r\n    \"unfilledtimeout\": {\r\n        \"buy\": 10,\r\n        \"sell\": 30\r\n    },\r\n    \"bid_strategy\": {\r\n        \"ask_last_balance\": 0.0,\r\n        \"use_order_book\": false,\r\n        \"order_book_top\": 1,\r\n        \"check_depth_of_market\": {\r\n            \"enabled\": false,\r\n            \"bids_to_ask_delta\": 1\r\n        }\r\n    },\r\n    \"ask_strategy\": {\r\n        \"use_order_book\": false,\r\n        \"order_book_min\": 1,\r\n        \"order_book_max\": 1,\r\n        \"use_sell_signal\": true,\r\n        \"sell_profit_only\": false,\r\n        \"ignore_roi_if_buy_signal\": false\r\n    },\r\n    \"exchange\": {\r\n        \"name\": \"binance\",\r\n        \"key\": \"MUEI9XWUexwPHR6zWu3o3VLZ0sY2eQ021BEXGBz2t5f77OX5XlwbONpSP2trQeKi\",\r\n        \"secret\": \"h2b0o7EMwf3qHFHEtTFPXkqHCF9BkCGKSil1qx4oUHL1alwVWejilCSTUkABWnRd\",\r\n        \"ccxt_config\": {\"enableRateLimit\": true},\r\n        \"ccxt_async_config\": {\r\n            \"enableRateLimit\": true,\r\n            \"rateLimit\": 200\r\n        },\r\n        \"pair_whitelist\": [\r\n            \"BTC\/USDT\",\r\n            \"1INCH\/USDT\",\r\n            \"AAVE\/USDT\",\r\n            \"ACM\/USDT\",\r\n            \"ADA\/USDT\",\r\n            \"AION\/USDT\",\r\n            \"AKRO\/USDT\",\r\n            \"ALGO\/USDT\",\r\n            \"ALICE\/USDT\",\r\n            \"ALPHA\/USDT\",\r\n            \"ANKR\/USDT\",\r\n            \"ANT\/USDT\",\r\n            \"ARDR\/USDT\",\r\n            \"ARPA\/USDT\",\r\n            \"ASR\/USDT\",\r\n            \"ATM\/USDT\",\r\n            \"ATOM\/USDT\",\r\n            \"AUDIO\/USDT\",\r\n            \"AUD\/USDT\",\r\n            \"AUTO\/USDT\",\r\n            \"AVA\/USDT\",\r\n            \"AVAX\/USDT\",\r\n            \"AXS\/USDT\",\r\n            \"BADGER\/USDT\",\r\n            \"BAL\/USDT\",\r\n            \"BAND\/USDT\",\r\n            \"BAT\/USDT\",\r\n            \"BCH\/USDT\",\r\n            \"BEAM\/USDT\",\r\n            \"BEL\/USDT\",\r\n            \"BLZ\/USDT\",\r\n            \"BNT\/USDT\",\r\n            \"BTCST\/USDT\",\r\n            \"BTT\/USDT\",\r\n            \"BUSD\/USDT\",\r\n            \"BZRX\/USDT\",\r\n            \"CAKE\/USDT\",\r\n            \"CELO\/USDT\",\r\n            \"CELR\/USDT\",\r\n            \"CFX\/USDT\",\r\n            \"CHR\/USDT\",\r\n            \"CHZ\/USDT\",\r\n            \"CKB\/USDT\",\r\n            \"COCOS\/USDT\",\r\n            \"COMP\/USDT\",\r\n            \"COS\/USDT\",\r\n            \"COTI\/USDT\",\r\n            \"CRV\/USDT\",\r\n            \"CTK\/USDT\",\r\n            \"CTSI\/USDT\",\r\n            \"CTXC\/USDT\",\r\n            \"CVC\/USDT\",\r\n            \"DASH\/USDT\",\r\n            \"DATA\/USDT\",\r\n            \"DCR\/USDT\",\r\n            \"DEGO\/USDT\",\r\n            \"DENT\/USDT\",\r\n            \"DGB\/USDT\",\r\n            \"DIA\/USDT\",\r\n            \"DNT\/USDT\",\r\n            \"DOCK\/USDT\",\r\n            \"DODO\/USDT\",\r\n            \"DOGE\/USDT\",\r\n            \"DOT\/USDT\",\r\n            \"DREP\/USDT\",\r\n            \"DUSK\/USDT\",\r\n            \"EGLD\/USDT\",\r\n            \"ENJ\/USDT\",\r\n            \"EOS\/USDT\",\r\n            \"EPS\/USDT\",\r\n            \"ETC\/USDT\",\r\n            \"EUR\/USDT\",\r\n            \"FET\/USDT\",\r\n            \"FIL\/USDT\",\r\n            \"FIO\/USDT\",\r\n            \"FIRO\/USDT\",\r\n            \"FIS\/USDT\",\r\n            \"FLM\/USDT\",\r\n            \"FTM\/USDT\",\r\n            \"FTT\/USDT\",\r\n            \"FUN\/USDT\",\r\n            \"GBP\/USDT\",\r\n            \"GRT\/USDT\",\r\n            \"GTO\/USDT\",\r\n            \"GXS\/USDT\",\r\n            \"HARD\/USDT\",\r\n            \"HBAR\/USDT\",\r\n            \"HIVE\/USDT\",\r\n            \"HNT\/USDT\",\r\n            \"HOT\/USDT\",\r\n            \"ICX\/USDT\",\r\n            \"INJ\/USDT\",\r\n            \"IOST\/USDT\",\r\n            \"IOTA\/USDT\",\r\n            \"IOTX\/USDT\",\r\n            \"IRIS\/USDT\",\r\n            \"JST\/USDT\",\r\n            \"JUV\/USDT\",\r\n            \"KAVA\/USDT\",\r\n            \"KEY\/USDT\",\r\n            \"KMD\/USDT\",\r\n            \"KNC\/USDT\",\r\n            \"KSM\/USDT\",\r\n            \"LINA\/USDT\",\r\n            \"LINK\/USDT\",\r\n            \"LIT\/USDT\",\r\n            \"LRC\/USDT\",\r\n            \"LSK\/USDT\",\r\n            \"LTC\/USDT\",\r\n            \"LTO\/USDT\",\r\n            \"LUNA\/USDT\",\r\n            \"MANA\/USDT\",\r\n            \"MATIC\/USDT\",\r\n            \"MBL\/USDT\",\r\n            \"MDT\/USDT\",\r\n            \"MFT\/USDT\",\r\n            \"MITH\/USDT\",\r\n            \"MKR\/USDT\",\r\n            \"MTL\/USDT\",\r\n            \"NANO\/USDT\",\r\n            \"NBS\/USDT\",\r\n            \"NEAR\/USDT\",\r\n            \"NEO\/USDT\",\r\n            \"NMR\/USDT\",\r\n            \"NULS\/USDT\",\r\n            \"OCEAN\/USDT\",\r\n            \"OGN\/USDT\",\r\n            \"OG\/USDT\",\r\n            \"OMG\/USDT\",\r\n            \"OM\/USDT\",\r\n            \"ONE\/USDT\",\r\n            \"ONG\/USDT\",\r\n            \"ONT\/USDT\",\r\n            \"ORN\/USDT\",\r\n            \"OXT\/USDT\",\r\n            \"PAXG\/USDT\",\r\n            \"PAX\/USDT\",\r\n            \"PERL\/USDT\",\r\n            \"PERP\/USDT\",\r\n            \"PNT\/USDT\",\r\n            \"POND\/USDT\",\r\n            \"PSG\/USDT\",\r\n            \"PUNDIX\/USDT\",\r\n            \"QTUM\/USDT\",\r\n            \"RAMP\/USDT\",\r\n            \"REEF\/USDT\",\r\n            \"REP\/USDT\",\r\n            \"RIF\/USDT\",\r\n            \"RLC\/USDT\",\r\n            \"ROSE\/USDT\",\r\n            \"RSR\/USDT\",\r\n            \"RUNE\/USDT\",\r\n            \"RVN\/USDT\",\r\n            \"SAND\/USDT\",\r\n            \"SC\/USDT\",\r\n            \"SFP\/USDT\",\r\n            \"SKL\/USDT\",\r\n            \"SNX\/USDT\",\r\n            \"SOL\/USDT\",\r\n            \"SRM\/USDT\",\r\n            \"STMX\/USDT\",\r\n            \"STORJ\/USDT\",\r\n            \"STPT\/USDT\",\r\n            \"STRAX\/USDT\",\r\n            \"STX\/USDT\",\r\n            \"SUN\/USDT\",\r\n            \"SUPER\/USDT\",\r\n            \"SUSD\/USDT\",\r\n            \"SUSHI\/USDT\",\r\n            \"SXP\/USDT\",\r\n            \"TCT\/USDT\",\r\n            \"TFUEL\/USDT\",\r\n            \"THETA\/USDT\",\r\n            \"TKO\/USDT\",\r\n            \"TOMO\/USDT\",\r\n            \"TRB\/USDT\",\r\n            \"TROY\/USDT\",\r\n            \"TRU\/USDT\",\r\n            \"TRX\/USDT\",\r\n            \"TUSD\/USDT\",\r\n            \"TWT\/USDT\",\r\n            \"UMA\/USDT\",\r\n            \"UNFI\/USDT\",\r\n            \"UNI\/USDT\",\r\n            \"USDC\/USDT\",\r\n            \"UTK\/USDT\",\r\n            \"VET\/USDT\",\r\n            \"VITE\/USDT\",\r\n            \"VTHO\/USDT\",\r\n            \"WAN\/USDT\",\r\n            \"WAVES\/USDT\",\r\n            \"WING\/USDT\",\r\n            \"WIN\/USDT\",\r\n            \"WNXM\/USDT\",\r\n            \"WRX\/USDT\",\r\n            \"WTC\/USDT\",\r\n            \"XEM\/USDT\",\r\n            \"XLM\/USDT\",\r\n            \"XMR\/USDT\",\r\n            \"XRP\/USDT\",\r\n            \"XTZ\/USDT\",\r\n            \"XVS\/USDT\",\r\n            \"YFII\/USDT\",\r\n            \"YFI\/USDT\",\r\n            \"ZEC\/USDT\",\r\n            \"ZEN\/USDT\",\r\n            \"ZIL\/USDT\",\r\n            \"ZRX\/USDT\",\r\n            \"BNB\/USDT\",\r\n            \"ETH\/USDT\",\r\n            \"REN\/USDT\",\r\n            \"NKN\/USDT\",\r\n            \"BTS\/USDT\"\r\n        ],\r\n        \"pair_blacklist\": [\r\n            \"BNB\/BTC\"\r\n        ]\r\n    },\r\n    \"pairlists\": [\r\n        {\"method\": \"StaticPairList\"}\r\n    ],\r\n    \"edge\": {\r\n        \"enabled\": false,\r\n        \"process_throttle_secs\": 3600,\r\n        \"calculate_since_number_of_days\": 7,\r\n        \"allowed_risk\": 0.01,\r\n        \"stoploss_range_min\": -0.0001,\r\n        \"stoploss_range_max\": -0.1,\r\n        \"stoploss_range_step\": -0.0001,\r\n        \"minimum_winrate\": 0.60,\r\n        \"minimum_expectancy\": 0.20,\r\n        \"min_trade_number\": 10,\r\n        \"max_trade_duration_minute\": 1440,\r\n        \"remove_pumps\": false\r\n    },\r\n    \"telegram\": {\r\n        \"enabled\": false,\r\n        \"token\": \"1777355043:AAFV61iE0FgXEkyZEbOdEYfsjldGI3cFehE\",\r\n        \"chat_id\": \"1574109528\"\r\n    },\r\n    \"api_server\": {\r\n        \"enabled\": true,\r\n        \"listen_ip_address\": \"0.0.0.0\",\r\n        \"listen_port\": 8080,\r\n        \"verbosity\": \"error\",\r\n        \"jwt_secret_key\": \"somethingrandom\",\r\n        \"CORS_origins\": [],\r\n        \"username\": \"freqtrader\",\r\n        \"password\": \"SuperSecurePassword\"\r\n    },\r\n    \"bot_name\": \"turantrade_bot\",\r\n    \"initial_state\": \"running\",\r\n    \"forcebuy_enable\": false,\r\n    \"internals\": {\r\n        \"process_throttle_secs\": 5\r\n    },\r\n    \"strategy\": \"PortfolioStrategy\"\r\n}"
}
```
## Delete a Trader
```POST``` /delete 
```json
{
	"user_id": "user_id",
	"trader_id": "trader_id"
}
```

# Building and Running

````shell
$ go build -o ./dist && eval $(minikube docker-env) && make dockerize 
# deploy postgres
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm install -f helm-values/postgres.yaml postgres bitnami/postgresql
# deploy model-manager for strategy deployment
$ kubectl apply -f deployments/model-manager
-------
## go to notebook & strategies folder
$ cd "wherever is the notebooks"
$ for i in $(ls *py); do k cp $i <pod_adı>:/tmp/strategies/$i; done;
$ for i in $(ls *pkl); do k cp $i <pod_adı>:/tmp/notebooks/$i; done;
## go back to repository, wait for postgres and deploy trader
$ kubectl apply -f deployments/trader

## Expose necessary service for testing
$ minikube service --url postgres-postgresql 
$ minikube service --url trader-provisioner
````



