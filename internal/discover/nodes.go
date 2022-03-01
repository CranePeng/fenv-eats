package discover

import (
	"context"
	"encoding/json"
	"fenv-eats/config"
	"fenv-eats/internal/pool"
	"fenv-eats/models"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientV3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type (
	Cluster struct {
		Client *clientV3.Client
	}
)

// 监听节点
func (c *Cluster) WatchNodes(ctx context.Context, id string) {
	log.Println("启动 etcd 监控节点任务")
	if Client == nil {
		NewClient()
	}
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

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			var node models.Node
			// 设置超时上下文，用来释放资源
			leaseCtx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
			switch event.Type {
			case mvccpb.PUT:
				// 如果需要添加或者更新节点，需要从etcd上取下来
				if err := json.Unmarshal(event.Kv.Value, &node); err != nil {
					log.Println(err)
				}
				result := seize(leaseCtx, c.Client, node.Id, id)
				if result {
					if err := node.CreateOrUpdate(); err != nil {
						log.Println(err)
					}
					log.Printf("节点：%s 注册成功", node.Id)
				}
				break
			case mvccpb.DELETE:
				if err := json.Unmarshal(event.PrevKv.Value, &node); err != nil {
					log.Println(err)
				}
				result := seize(leaseCtx, c.Client, node.Id, id)
				if result {
					node.Status = models.OFFLINE
					if err := node.Update(); err != nil {
						log.Println(err)
					}
					log.Printf("节点：%s 离线", node.Id)
				}
				break
			}
			// 处理完或者超时后要取消上下文数据
			cancelFunc()
		}
	}

}

// 抢用于更新节点信息的锁(分布式锁)
func seize(ctx context.Context, client *clientV3.Client, key, val string) bool {
	res, err := client.Grant(ctx, 5)
	if err != nil {
		log.Println(err)
		return false
	}
	// 持久化租约
	resc, err := client.KeepAlive(ctx, res.ID)
	if err != nil {
		log.Println(err)
		return false
	}
	antsPoolErr := pool.GetWorkerInstance().Work(func() {
		leaseWatch(ctx, res, resc)
	})
	if antsPoolErr != nil {
		log.Println(antsPoolErr)
		return false
	}
	lockey := fmt.Sprintf("%s/%s/", config.Conf.Locker, key)
	// 事务操作获取锁
	txn := Client.Txn(ctx)
	txn.If(clientV3.Compare(clientV3.CreateRevision(lockey), "=", 0)).
		Then(clientV3.OpPut(lockey, val, clientV3.WithLease(res.ID))).
		Else(clientV3.OpGet(lockey))
	txnResp, err := txn.Commit()
	if err != nil {
		log.Println(err)
		return false
	}

	if !txnResp.Succeeded {
		return false
	} else {
		return true
	}
}

// 监听租约
func leaseWatch(ctx context.Context, res *clientV3.LeaseGrantResponse, resc <-chan *clientV3.LeaseKeepAliveResponse) {
LOOP:
	for {
		select {
		case kresp := <-resc:
			if kresp != nil {
				log.Println("续租成功，LeaseID: ", kresp.ID)
			} else if resc == nil {
				log.Println("续租失败")
				break LOOP
			}
		case <-ctx.Done():
			if _, err := Client.Revoke(context.TODO(), res.ID); err != nil {
				log.Println(err)
			}
			break LOOP
		}
	}
	log.Println("leaseWatch done!")
}
