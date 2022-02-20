package models

type PasswordResets struct {
	Email string `gorm:"not null; index;type: VARCHAR(255)"`
	Token string `gorm:"not null; type:VARCHAR(255)"`
	CommonColumn
}

// 定义模型的数据表名称
func (resets *PasswordResets) TableName() string {
	return "password_resets"
}
