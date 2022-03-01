package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type (
	Meta struct {
		Page  int `json:"page"`  // 当前页
		Limit int `json:"limit"` // 每页数量
		Total int `json:"total"` // 总数
	}
	Response struct {
		// 业务编码,200成功状态下返回Data，其他错误状态返回ErrorCode,errorCode和message出现时，data是空
		Code    int         `json:"code"`    // 自定义错误码
		Data    interface{} `json:"data"`    // 数据
		Message string      `json:"message"` // 信息
		Meta    *Meta       `json:"meta,omitempty"`
	}
	Payload map[string]interface{}
)

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, payload map[string]interface{}) {
	data := payload["data"]
	meta, ok := payload["meta"]
	if ok {
		c.JSON(http.StatusOK, Response{
			http.StatusOK,
			data,
			"ok",
			reflect.ValueOf(meta).Interface().(*Meta),
		})
	}
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "ok",
	})

}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    errorCode,
		Data:    nil,
		Message: msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, errorCode int, msg string) {
	Fail(c, errorCode, msg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, Errors.BusinessError.ErrorCode, msg)
}
