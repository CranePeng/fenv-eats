package node

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller struct {
	}

	CreateRequest struct {
		Name   string `json:"name" validate:"required"`
		Remark string `json:"remark"`
	}

	UpdateRequest struct {
		Name   string `json:"name" validate:"required"`
		Remark string `json:"remark"`
	}
)

// 获取节点列表
func (instance *Controller) GetList(c *gin.Context) {

}

// 创建节点
func (instance *Controller) CreateNode(c *gin.Context) {

}

// 修改节点信息
func (instance *Controller) ModifyNode(c *gin.Context) {

}

// 获取节点关联的流水线
func (instance *Controller) GetPipelines(c *gin.Context) {

}

// 删除节点
func (instance *Controller) DeleteNode(c *gin.Context) {

}

// 关联流水线
func (instance *Controller) RelateWithPipeline(c *gin.Context) {

}

// 解绑流水线
func (instance *Controller) DisassociatePipeline(c *gin.Context) {

}
