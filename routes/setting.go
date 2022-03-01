package routes

import (
	"fenv-eats/handlers/global"
	"fenv-eats/handlers/setting"
	"github.com/gin-gonic/gin"
)

func registerSetting(e *gin.Engine) {
	handler := setting.Controller{}
	// 优化一下。这里用分组管理，便于使用中间件
	authorized := e.Group("/setting")
	authorized.Use(global.AuthRequired)
	{
		// 获取通知配置信息
		authorized.POST("/getNotification", handler.GetNotification)
		// 更新服务配置
		authorized.POST("/modifyServerConfig", handler.ModifyServerConfig)
		// 发送邮件
		authorized.POST("/sendEMail", handler.SendEMail)

	}

}
