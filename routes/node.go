package routes

import (
	"fenv-eats/handlers/global"
	"fenv-eats/handlers/node"
	"github.com/gin-gonic/gin"
)

func registerNode(e *gin.Engine) {
	handler := node.Controller{}
	// 优化一下。这里用分组管理，便于使用中间件
	authorized := e.Group("/node")
	authorized.Use(global.AuthRequired)
	{
		// 加载列表
		authorized.POST("/getList", handler.GetList)
		// 删除节点
		authorized.POST("/deleteNode", handler.DeleteNode)
		// 创建节点
		authorized.POST("/createNode", handler.CreateNode)
		// 修改节点
		authorized.POST("/modifyNode", handler.ModifyNode)
		// 关联流水线
		authorized.POST("/relateWithPipeline", handler.RelateWithPipeline)
		// 解绑流水线
		authorized.POST("/disassociatePipeline", handler.DisassociatePipeline)
	}

}
