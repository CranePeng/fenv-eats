package discover

import (
	"fenv-eats/config"
	"fenv-eats/internal/service"
	clientV3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

type (
	Service struct {
		instance *service.Instance
		leaseID  clientV3.LeaseID
		close    chan struct{}
		wg       sync.WaitGroup
	}
)

var (
	err    error
	Client *clientV3.Client
)

// New ETCD V3 Client
func NewClient() {
	etcdPoints := config.Conf.Etcd.EndPoints
	if config.DockerConf.EtcdIp != nil && len(config.DockerConf.EtcdIp) > 0 {
		etcdPoints = config.DockerConf.EtcdIp
	}
	if Client, err = clientV3.New(clientV3.Config{
		Endpoints:   etcdPoints,
		DialTimeout: 10 * time.Second,
	}); err != nil {
		log.Println(err)
	}
}
