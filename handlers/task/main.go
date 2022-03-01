package task

import (
	"fenv-eats/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	Controller struct {
	}

	UpdateRequest struct {
		Name        string `json:"name" validate:"required"`
		Content     string `json:"content" validate:"required"`
		Description string `json:"description"`
	}
)

var (
	validate = validator.New()
)

// 获取任务列表
func (instance *Controller) GetTaskList(ctx *gin.Context) {
	response.Fail(ctx, 505, "数据使用场景有误")
}

// 创建任务
func (instance *Controller) CreateTask(ctx *gin.Context) {

}

// 更新任务
func (instance *Controller) ModifyTask(ctx *gin.Context) {

}

// 删除任务
func (instance *Controller) DeleteTask(ctx *gin.Context) {

}
