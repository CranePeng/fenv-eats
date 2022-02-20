package discover

import (
	"context"
	"fenv-eats/config"
	clientV3 "go.etcd.io/etcd/client/v3"
)

var ServiceCluster *Cluster

type (
	Cluster struct {
		client *clientV3.Client
	}
)

// 监听节点
func (c *Cluster) WatchNodes(ctx context.Context, key string) {
	if Client == nil {
		NewClient()
	}
	c.client = Client
	var curRevision int64 = 0
	for {
		rangeResp, err := Client.Get(context.TODO(), config.Conf.Etcd.Service, clientV3.WithPrefix())

		if err != nil {
			continue
		}
		// 从当前版本开始订阅
		curRevision = rangeResp.Header.Revision + 1
		break
	}
	// 监听通道
	watchChan := Client.Watch(ctx, config.Conf.Etcd.Service, clientV3.WithPrefix(), clientV3.WithRev(curRevision), clientV3.WithPrevKV())

}

// 抢用于更新节点信息的锁
func seize(ctx context.Context, key, value string) bool {

}
