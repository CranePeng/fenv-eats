package organization

import (
	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
	}

	CreateRequest struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Pass    string `json:"pass" validate:"required"`
		Confirm string `json:"confirm" validate:"eqfield=Pass"`
		Manager bool   `json:"manager"`
	}

	UpdateRequest struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Manager bool   `json:"manager"`
	}
)

// 获取用户列表
func (instance *UserController) Get(c *gin.Context) {

}

// 创建用户
func (instance *UserController) Post(c *gin.Context) {

}

// 修改用户信息
func (instance *UserController) PutBy(c *gin.Context) {

}

// 删除用户
func (instance *UserController) DeleteBy(c *gin.Context) {

}
