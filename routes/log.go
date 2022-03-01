package routes

import (
	"fenv-eats/handlers/global"
	"fenv-eats/handlers/log"
	"github.com/gin-gonic/gin"
)

func registerLog(e *gin.Engine) {
	handler := log.Controller{}
	// 优化一下。这里用分组管理，便于使用中间件
	authorized := e.Group("/log")
	authorized.Use(global.AuthRequired)
	{
		//  获取日志数据
		authorized.POST("/getList", handler.GetList)

	}

}
