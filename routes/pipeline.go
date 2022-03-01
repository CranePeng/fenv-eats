package routes

import (
	"fenv-eats/handlers/global"
	"fenv-eats/handlers/pipeline"
	"github.com/gin-gonic/gin"
)

func registerPipeline(e *gin.Engine) {
	handler := pipeline.Controller{}
	// 优化一下。这里用分组管理，便于使用中间件
	pipeline := e.Group("/pipeline")
	pipeline.Use(global.AuthRequired)
	{
		// 获取列表
		pipeline.POST("/getList", handler.GetList)
		// 创建流水线
		pipeline.POST("/create", handler.Create)
		// 流水线绑定任务
		pipeline.POST("/relateWithTask", handler.RelateWithTask)
		// 创建强杀指令
		pipeline.POST("/createKiller", handler.CreateKiller)
		// 删除流水线
		pipeline.POST("/deletePipeline", handler.DeletePipeline)
		// 根据流水线获取节点数据
		pipeline.POST("/getNodes", handler.GetNodes)
		// 关联流水线到节点
		pipeline.POST("/relateToNodes", handler.RelateToNodes)
		// 修改流水线绑定关系
		pipeline.POST("/modifyRelationship", handler.ModifyRelationship)
		// 更新流水线
		pipeline.POST("/modifyPipeline", handler.ModifyPipeline)
		// 同步流水线数据到etcd
		pipeline.POST("/syncPipelineToETCD", handler.SyncPipelineToETCD)
		// 获取流水线绑定的任务
		pipeline.POST("/getTasksBy", handler.GetTasksBy)
		// 根据拖动排序
		pipeline.POST("/sortByDrag", handler.SortByDrag)
		// 流水线解绑任务
		pipeline.POST("/disassociateTask", handler.DisassociateTask)

	}

}
