package models

type Log struct {
	Id        int64  `json:"id" gorm:"primary_key; AUTO_INCREMENT; comment:'ID'; type:BIGINT(20)"`
	UserId    string `json:"user_id" gorm:"not null; comment:'用户ID'; index ;type:CHAR(36)"`
	Operation string `json:"operation" gorm:"not null; comment:'操作';type: VARCHAR(255)"`
	Result    string `json:"result" gorm:"null; comment:'结果'; type: LONGTEXT"`
	CommonColumn
}

// 定义日志表名称
func (log *Log) TableName() string {
	return "logs"
}
