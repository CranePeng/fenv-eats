package models

import (
	"time"
)

type PipelineTaskPivot struct {
	Id          string    `json:"id" gorm:"not null ;primary_key; comment:'ID';type: CHAR(36)"`
	PipelineId  string    `json:"pipeline_id" validate:"required,uuid4" gorm:"not null; comment:'ID'; index;type: CHAR(36)"`
	TaskId      string    `json:"task_id" validate:"required,uuid4" gorm:"not null; comment:'ID'; index;type: CHAR(36)"`
	Step        int       `json:"step" validate:"numeric" gorm:"not null; comment:'步骤';type: SMALLINT(5)"`
	Timeout     int       `json:"timeout" validate:"numeric" gorm:"not null; default 0; comment:'超时时间';type: INT(10)"`
	Interval    int       `json:"interval" validate:"numeric" gorm:"not null; default 0; comment:'间隔时间';type: INT(10)"`
	Retries     int       `json:"retries" validate:"numeric" gorm:"not null; default 0; comment:'重试次数';type: TINYINT(3)"`
	Directory   string    `json:"directory" validate:"omitempty" gorm:"null; comment:'工作目录';type: VARCHAR(255)"`
	User        string    `json:"user" validate:"omitempty" gorm:"null; comment:'运行用户';type: VARCHAR(255)"`
	Environment string    `json:"environment" validate:"omitempty" gorm:"null; comment:'环境变量';type: VARCHAR(255)"`
	Dependence  string    `json:"dependence" validate:"required" gorm:"not null; default 'strong'; comment:'依赖';type: VARCHAR(255)"`
	CreatedAt   time.Time `json:"created_at" validate:"-" gorm:"not null; comment:'创建于';type: DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" validate:"-" gorm:"not null; comment:'更新于';type: DATETIME"`
	Task        *Task     `json:"task" validate:"-" gorm:"-"`
}

// 定义模型的数据表名称
func (pivot *PipelineTaskPivot) TableName() string {
	return "pipeline_task_pivot"
}
