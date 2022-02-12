package models

import (
	"encoding/json"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id" gorm:"not null; primary_key; comment:'用户ID';type:CHAR(36)"`
	Name      string    `json:"name" gorm:"not null; comment:'姓名'; type:VARCHAR(255)"`
	Email     string    `json:"email" gorm:"not null; comment:'邮箱'; unique_index:idx_user_email; type:VARCHAR(255)"`
	Password  string    `json:"-" gorm:"not null; comment:'密码'; type:VARCHAR(255)"`
	Manager   bool      `json:"manager" gorm:"not null; default 0; comment:'管理员'; type:TINYINT(1)"`
	CreatedAt time.Time `json:"created_time" gorm:"column:created_at;not null; comment:'创建于'; type:DATETIME"`
	UpdatedAt time.Time `json:"updated_time" gorm:"column:updated_at;not null; comment:'更新于'; type:DATETIME"`
}

// Define table name
func (user *User) TableName() string {
	return "users"
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

// 更新用户信息
func (user *User) Update() error {
	err := Engine.Model(user).Updates(user).Error
	return err
}

func (user *User) Store() error {
	err := Engine.Create(user).Error
	return err
}

func (user *User) Save() error {
	err := Engine.Model(user).Updates(user).Error
	return err
}

func (user *User) ModifyEmail(email string) (*User, error) {
	user.Email = email
	err := user.Save()
	return user, err
}

// 序列化
func (user *User) ToString() (string, error) {
	result, err := json.Marshal(user)
	return string(result), err
}
