package pipeline

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	Controller struct {
	}
	BindNodeRequest struct {
		PipelineId string   `json:"pipeline_id" validate:"required,uuid4"`
		NodesId    []string `json:"nodes_id" validate:"required"`
	}
	KillPipelineRequest struct {
		PipelineId string `json:"pipeline_id" validate:"required,uuid4"`
	}
	PutStepsRequest struct {
		PipelineId string `json:"pipeline_id" validate:"required,uuid4"`
		Origin     int    `json:"origin" validate:"numeric"`
		Current    int    `json:"current" validate:"numeric"`
	}
)

var (
	validate = validator.New()
)

// 获取流水线列表
func (instance *Controller) GetList(c *gin.Context) {

}

// 创建流水线
func (instance *Controller) Create(c *gin.Context) {

}

// 更新流水线
func (instance *Controller) ModifyPipeline(c *gin.Context) {

}

// 删除流水线
func (instance *Controller) DeletePipeline(c *gin.Context) {

}

// 获取流水线绑定的节点
func (instance *Controller) GetNodes(c *gin.Context) {

}

// 绑定流水线到节点
func (instance *Controller) RelateToNodes(c *gin.Context) {

}

// 获取流水线绑定的任务
func (instance *Controller) GetTasksBy(c *gin.Context) {

}

// 根据拖动顺序排序数据
func (instance *Controller) SortByDrag(c *gin.Context) {

}

// 绑定任务到流水线
func (instance *Controller) RelateWithTask(c *gin.Context) {

}

// 修改绑定关系
func (instance *Controller) ModifyRelationship(c *gin.Context) {

}

// 从流水线解绑任务
func (instance *Controller) DisassociateTask(c *gin.Context) {

}

// 同步流水线数据到 ETCD
func (instance *Controller) SyncPipelineToETCD(c *gin.Context) {

}

// 创建强杀指令
func (instance *Controller) CreateKiller(c *gin.Context) {

}
