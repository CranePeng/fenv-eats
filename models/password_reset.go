package models

import (
	"time"
)

type PasswordResets struct {
	Email     string    `gorm:"not null; index;type: VARCHAR(255)"`
	Token     string    `gorm:"not null; type:VARCHAR(255)"`
	CreatedAt time.Time `gorm:"not null; comment:'创建于';type: DATETIME"`
}

// 定义模型的数据表名称
func (resets *PasswordResets) TableName() string {
	return "password_resets"
}
