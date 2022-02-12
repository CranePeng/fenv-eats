package pool

import (
	"fenv-eats/log"
	"github.com/panjf2000/ants/v2"
	"runtime"
)

/*
	协程池采用 ants : https://github.com/panjf2000/ants
*/

var (
	workPool *ants.Pool
)

func Init() {
	workPool, err := ants.NewPool(runtime.NumCPU() * 20)
	if err != nil {
		log.ErrorF("初始化协程池失败：%v", err)
	}
	// 记录协程池地址
	log.InfoF("初始化协程池成功，协程池:%P", workPool)
}

func Work(w func()) error {
	return workPool.Submit(w)
}
