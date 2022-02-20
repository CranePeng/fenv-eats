package models

import (
	"time"
)

type TaskRecords struct {
	Id               int64     `json:"id" gorm:"primary_key; auto_increment; comment:'ID';type: BIGINT(20)"`
	PipelineRecordId string    `json:"pipeline_record_id" gorm:"not null; comment:'流水线记录ID'; index;type: CHAR(36)"`
	TaskId           string    `json:"task_id" gorm:"not null; comment:'任务ID'; index; type:CHAR(36)"`
	NodeId           string    `json:"node_id" gorm:"not null; comment:'节点ID'; index; type:CHAR(36)"`
	TaskName         string    `json:"task_name" gorm:"not null; comment:'任务名称'; type:VARCHAR(255)"`
	WorkerName       string    `json:"worker_name" gorm:"not null; comment:'节点名称';type: VARCHAR(255)"`
	Content          string    `json:"content" gorm:"not null; comment:'执行内容';type: TEXT"`
	Mode             string    `json:"mode" gorm:"not null; comment:'执行方式';type: VARCHAR(255)"`
	Timeout          int       `json:"timeout" gorm:"not null; default 0; comment:'超时时间'; type:INT(10)"`
	Retries          int       `json:"retries" gorm:"not null; default 0; comment:'重试次数'; type:TINYINT(3)"`
	Status           string    `json:"status" gorm:"not null; default 'finished'; comment:'状态'; type:VARCHAR(255)"`
	Result           string    `json:"result" gorm:"not null; comment:'执行结果'; type:TEXT"`
	Duration         int64     `json:"duration" gorm:"not null; comment:'持续时间'; type:INT(10)"`
	BeginWith        time.Time `json:"begin_with" gorm:"not null; comment:'开始于'; type:DATETIME"`
	FinishWith       time.Time `json:"finish_with" gorm:"not null; comment:'结束于'; type:DATETIME"`
	CommonColumn
}

// 定义模型的数据表名称
func (records *TaskRecords) TableName() string {
	return "task_records"
}
