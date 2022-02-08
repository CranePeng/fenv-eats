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
	if Client, err = clientV3.New(clientV3.Config{
		Endpoints:   config.Conf.Etcd.EndPoints,
		DialTimeout: 10 * time.Second,
	}); err != nil {
		log.Println(err)
	}
}
