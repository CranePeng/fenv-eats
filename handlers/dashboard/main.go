package dashboard

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller struct{}
)

// 获取节点数据
func (instance *Controller) GetNodes(c *gin.Context) {

}

// 获取正在调度的流水线数量
func (instance *Controller) GetPipelines(c *gin.Context) {

}

// 获取流水线失败次数
func (instance *Controller) GetFailCount(c *gin.Context) {

}
