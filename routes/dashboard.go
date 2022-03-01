package routes

import (
	"fenv-eats/handlers/dashboard"
	"fenv-eats/handlers/global"
	"github.com/gin-gonic/gin"
)

func registerDashboard(e *gin.Engine) {
	handler := dashboard.Controller{}
	// 优化一下。这里用分组管理，便于使用中间件
	authorized := e.Group("/dashboard")
	authorized.Use(global.AuthRequired)
	{
		// 登录
		authorized.POST("/getNodes", handler.GetNodes)
		// 注册
		authorized.POST("/getPipelines", handler.GetPipelines)
		// 登出
		authorized.POST("/getFailCount", handler.GetFailCount)
	}

}
