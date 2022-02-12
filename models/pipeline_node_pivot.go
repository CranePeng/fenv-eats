package models

import (
	"time"
)

type PipelineNodePivot struct {
	Id         string    `json:"id" gorm:"not null; primary_key; comment:'ID';type: CHAR(36)"`
	PipelineId string    `json:"pipeline_id" validate:"required,uuid4" gorm:"not null; index; comment:'流水线ID';type: CHAR(36)"`
	NodeId     string    `json:"node_id" validate:"required,uuid4" gorm:"not null; index; comment:'节点ID'; type:CHAR(36)"`
	CreatedAt  time.Time `json:"created_at" validate:"-" gorm:"not null; comment:'创建于';type: DATETIME"`
	Pipeline   *Pipeline `json:"pipeline" gorm:"-"`
}

// 定义模型的数据表名称
func (pivot *PipelineNodePivot) TableName() string {
	return "pipeline_node_pivot"
}
