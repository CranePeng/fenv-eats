package discover

import (
	"context"
	"encoding/json"
	"fenv-eats/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

// Get config from etcd
func GetConf(key string) {
	if res, err := Client.Get(context.TODO(), key, clientv3.WithFirstKey()...); err != nil {
		log.Fatal(err)
	} else if res != nil {
		if res.Kvs == nil {
			log.Fatal("config key not exist")
		}
		if err := json.Unmarshal(res.Kvs[0].Value, &config.Conf); err != nil {
			log.Fatal(err)
		}
	}
}
