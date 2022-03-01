package models

const (
	MODE_SHELL = "shell"
	MODE_HTTP  = "http"
	MODE_MAIL  = "mail"
	MODE_HOOK  = "hook"
)

// 任务模型
type Task struct {
	Id          string `json:"id" validate:"-" gorm:"not null; primary_key; comment:'用户ID'; type: CHAR(36)"`
	Name        string `json:"name" validate:"required" gorm:"not null; comment:'名称'; type VARCHAR(255)"`
	Mode        string `json:"mode" validate:"required" gorm:"not null; default: 'shell'; comment:'任务模式'; type:VARCHAR(32)"`
	Url         string `json:"url" validate:"omitempty" gorm:"comment:'请求URL'; type:VARCHAR(255)"`
	Method      string `json:"method" validate:"omitempty" gorm:"comment:'任务模式'; type:VARCHAR(255)"`
	Content     string `json:"content" validate:"omitempty" gorm:"comment:'内容'; type: TEXT"`
	Description string `json:"description" validate:"-" gorm:"comment:'描述'; type: VARCHAR(255)"`
	CommonColumn
}

// 定义模型的数据表名称
func (task *Task) TableName() string {
	return "tasks"
}

// 创建
func (task *Task) Create() error {
	err := Engine.Create(task).Error
	return err
}

// 更新
func (task *Task) Update() int {
	i := Engine.Model(task).Updates(task).RowsAffected
	return int(i)
}

// 软删除
func (task *Task) Delete() error {
	err := Engine.Delete(task).Error
	return err
}
