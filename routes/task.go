package routes

import (
	"fenv-eats/handlers/global"
	"fenv-eats/handlers/task"
	"github.com/gin-gonic/gin"
)

func registerTask(e *gin.Engine) {
	handler := task.Controller{}
	// 优化一下。这里用分组管理，便于使用中间件
	authorized := e.Group("/task")
	authorized.Use(global.AuthRequired)
	{
		// 获取任务列表
		authorized.POST("/getTaskList", handler.GetTaskList)
		// 创建任务
		authorized.POST("/createTask", handler.CreateTask)
		// 更新任务
		authorized.POST("/modifyTask", handler.ModifyTask)
		// 删除任务
		authorized.POST("/deleteTask", handler.DeleteTask)

	}
}
