package setting

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller struct{}
)

// 获取通知配置信息
func (instance *Controller) GetNotification(c *gin.Context) {

}

// 测试发送邮件功能
func (instance *Controller) SendEMail(c *gin.Context) {

}

// 更新服务配置
func (instance *Controller) ModifyServerConfig(c *gin.Context) {

}
