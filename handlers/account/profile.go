package account

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller struct {
	}
	Profile struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		TeamId  string `json:"team_id"`
		Manager bool   `json:"manager"`
	}
)

// 获取用户信息
func (instance *Controller) loadAccount(ctx *gin.Context) {

}

// 修改用户信息
func (instance *Controller) UpdateAccount(ctx *gin.Context) {

}
