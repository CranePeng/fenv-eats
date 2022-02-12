package models

import (
	"time"
)

const (
	MODE_SHELL = "shell"
	MODE_HTTP  = "http"
	MODE_MAIL  = "mail"
	MODE_HOOK  = "hook"
)

// 任务模型
type Task struct {
	Id          string    `json:"id" validate:"-" gorm:"not null; primary_key; comment:'用户ID'; type: CHAR(36)"`
	Name        string    `json:"name" validate:"required" gorm:"not null; comment:'名称'; type VARCHAR(255)"`
	Mode        string    `json:"mode" validate:"required" gorm:"not null; default: 'shell'; comment:'任务模式'; type:VARCHAR(32)"`
	Url         string    `json:"url" validate:"omitempty" gorm:"comment:'请求URL'; type:VARCHAR(255)"`
	Method      string    `json:"method" validate:"omitempty" gorm:"comment:'任务模式'; type:VARCHAR(255)"`
	Content     string    `json:"content" validate:"omitempty" gorm:"comment:'内容'; type: TEXT"`
	Description string    `json:"description" validate:"-" gorm:"comment:'描述'; type: VARCHAR(255)"`
	CreatedAt   time.Time `json:"created_at" validate:"-" gorm:"not null; comment:'创建于' type:DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" validate:"-" gorm:"not null; comment:'更新于' type DATETIME"`
}

// 定义模型的数据表名称
func (task *Task) TableName() string {
	return "tasks"
}
