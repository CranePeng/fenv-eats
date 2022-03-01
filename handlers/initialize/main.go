package initialize

import (
	"fenv-eats/config"
	"github.com/gin-gonic/gin"
)

type (
	Controller struct {
	}
	PostRequest struct {
		Etcd     config.Etcd     `json:"etcd"`
		Database config.Database `json:"database"`
		User     config.User     `json:"user"`
		Auth     config.Auth     `json:"auth"`
	}
)

// 获取系统信息
func (instance *Controller) Get(c *gin.Context) {
}

// 创建服务配置
func (instance *Controller) Post(c *gin.Context) {

}

// 生成 JWT Secret
func (instance *Controller) GetSecret(c *gin.Context) {

}

// 验证数据库是否存在
func (instance *Controller) GetDatabase(c *gin.Context) {

}
