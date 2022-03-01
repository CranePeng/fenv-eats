package routes

import (
	"fenv-eats/handlers/global"
	"fenv-eats/handlers/user"
	"github.com/gin-gonic/gin"
)

// 用户模块路由
func registerUser(e *gin.Engine) {
	handler := user.NewHandler()
	// 优化一下。这里用分组管理，便于使用中间件
	authorized := e.Group("/user")
	authorized.Use(global.AuthRequired)
	{
		// 登录
		authorized.POST("/login", handler.Login)
		// 注册
		authorized.POST("/logout", handler.Logout)
		// 登出
		authorized.POST("/registry", handler.Registry)

	}

}
