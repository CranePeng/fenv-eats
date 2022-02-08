package models

import "fenv-eats/internal/utils"

const (
	ONLINE  = "online"
	OFFLINE = "offline"
	MASTER  = "master"
	WORKER  = "worker"
)

type (
	Node struct {
		Id          string               `json:"id" gorm:"not null; primary_key; comment:'用户ID'; type:CHAR(36)"`
		Name        string               `json:"name" gorm:"not null; comment:'名称'; type:VARCHAR(255)"`
		Host        string               `json:"host" gorm:"not null; comment:'主机地址'; type:VARCHAR(255)"`
		Port        int                  `json:"port" gorm:"not null; comment:'端口';type: SMALLINT(5)"`
		Mode        string               `json:"mode" gorm:"not null; comment:'节点类型' CHAR(6)"`
		Status      string               `json:"status" gorm:"not null; default:'connected' comment:'状态'; type:VARCHAR(255)"` // 状态
		Version     string               `json:"version" gorm:"not null; comment:'版本'; type:VARCHAR(255)"`                    // 版本
		Description string               `json:"description" gorm:"comment:'描述'; type:VARCHAR(255)"`                          // 描述信息
		CreatedAt   utils.Time           `json:"created_at" gorm:"not null; comment:'创建于'; type:DATETIME"`                    // 创建于
		UpdatedAt   utils.Time           `json:"updated_at" gorm:"not null; comment:'更新于'; type:DATETIME"`                    // 更新于
		Pipelines   []*PipelineNodePivot `json:"pipelines" gorm:"-"`                                                          // 关联的流水线
	}
)

// 定义模型的数据表名称
func (node *Node) TableName() string {
	return "nodes"
}