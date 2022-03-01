package user

import (
	"github.com/gin-gonic/gin"
	"log"
)

type (
	Controller interface {
		// 登录
		Login(c *gin.Context)
		// 登出
		Logout(c *gin.Context)
		// 注册
		Registry(c *gin.Context)
		// 修改用户信息
		UpdateUser(c *gin.Context)
	}
	Handler struct {
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

func NewHandler() Controller {
	return &Handler{}
}

func (h *Handler) Login(c *gin.Context) {
	log.Println("测试登录接口")
}
func (h *Handler) Logout(c *gin.Context) {

}
func (h *Handler) Registry(c *gin.Context) {

}
func (h *Handler) UpdateUser(c *gin.Context) {

}
