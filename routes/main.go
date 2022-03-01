package routes

import (
	"fenv-eats/handlers/global"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

func Init(e *gin.Engine) *gin.Engine {
	include(
		registerUser,
		registerNode,
		registerTask,
		registerLog,
		registerSetting,
		registerPipeline,
		registerDashboard,
	)
	return register(e)
}

// Include 注册app的路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// 路由注册
func register(engine *gin.Engine) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(global.Cors())
	for _, opt := range options {
		opt(engine)
	}
	return engine

}
