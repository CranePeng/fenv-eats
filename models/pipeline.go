package models

import (
	"github.com/gorhill/cronexpr"
	"time"
)

// 流水线模型
type Pipeline struct {
	Id           string               `json:"id" validate:"-" gorm:"not null; primaryKey; comment:'ID';type:CHAR(36)"`
	Name         string               `json:"name" validate:"required" gorm:"not null; comment:'名称'; type:VARCHAR(255)"`
	Description  string               `json:"description" validate:"-" gorm:"not null; comment:'描述'; type:VARCHAR(255)"`
	Spec         string               `json:"spec" validate:"required" gorm:"not null; comment:'定时器'; type:CHAR(64)"`
	Status       int                  `json:"status" validate:"numeric" gorm:"not null; default 0; comment:'状态'; type:TINYINT(1)"`
	Finished     string               `json:"finished" validate:"omitempty,uuid4" gorm:"null; comment:'成功时执行'; type:CHAR(36)"`
	Failed       string               `json:"failed" validate:"omitempty,uuid4" gorm:"null; comment:'失败时执行'; type:CHAR(36)"`
	Overlap      int                  `json:"overlap" validate:"numeric" gorm:"not null; default 0 comment:'重复执行'; type:TINYINT(1)"`
	Nodes        []string             `json:"nodes" gorm:"-"`
	Steps        []*PipelineTaskPivot `json:"steps" gorm:"-"`
	Expression   *cronexpr.Expression `json:"-" gorm:"-"`
	NextTime     time.Time            `json:"-" gorm:"-"`
	FinishedTask *Task                `json:"finished_task,omitempty" gorm:"-"`
	FailedTask   *Task                `json:"failed_task,omitempty" gorm:"-"`
	CommonColumn
}

// 定义模型的数据表名称
func (pipeline *Pipeline) TableName() string {
	return "pipelines"
}
