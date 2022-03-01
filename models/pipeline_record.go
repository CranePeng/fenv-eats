package models

import (
	"time"
)

type (
	// 流水线调度记录模型
	PipelineRecords struct {
		Id         string         `json:"id" gorm:"not null; primary_key; comment:'ID';type: CHAR(36)"`
		PipelineId string         `json:"pipeline_id" gorm:"not null; comment:'流水线ID'; index;type: CHAR(36)"`
		NodeId     string         `json:"node_id" gorm:"not null; comment:'节点ID'; index; type:CHAR(36)"`
		WorkerName string         `json:"worker_name" gorm:"not null; comment:'节点名称';type: VARCHAR(255)"`
		Spec       string         `json:"spec" gorm:"comment:'定时器';type: CHAR(64)"`
		Status     int            `json:"status" gorm:"not null; default 1; comment:'状态';type: TINYINT(1)"`
		Duration   int64          `json:"duration" gorm:"not null; comment:'持续时间';type: INT(10)"`
		BeginWith  time.Time      `json:"begin_with" gorm:"not null; comment:'开始于';type: DATETIME"`
		FinishWith time.Time      `json:"finish_with" gorm:"not null; comment:'结束于';type: DATETIME"`
		Steps      []*TaskRecords `json:"steps" gorm:"-"`
		CommonColumn
	}
	// 流水线执行结果
	Result struct {
		Pipeline *PipelineRecords // 流水线执行记录
		Steps    []*TaskRecords   // 流水线中任务执行记录
	}
)

// 定义模型的数据表名称
func (records *PipelineRecords) TableName() string {
	return "pipeline_records"
}

// 创建
func (records *PipelineRecords) Create() error {
	err := Engine.Create(records).Error
	return err
}

// 更新
func (records *PipelineRecords) Update() int {
	i := Engine.Model(records).Updates(records).RowsAffected
	return int(i)
}

// 软删除
func (records *PipelineRecords) Delete() error {
	err := Engine.Delete(records).Error
	return err
}
