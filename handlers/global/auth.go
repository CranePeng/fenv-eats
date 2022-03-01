package global

import (
	"fenv-eats/internal/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// 最外层认证请求中间件
func AuthRequired(c *gin.Context) {
	t := time.Now()
	//过滤是否验证token， login结构直接放行，这里为了简单起见，直接判断路径中是否带login，携带login直接放行
	if strings.Contains(c.Request.RequestURI, "login") || strings.Contains(c.Request.RequestURI, "signIn") {
		return
	}
	token := c.Request.Header.Get("token")
	if len(token) < 1 {
		response.Fail(c, 700, "请求未携带token，无权限访问")
		c.Abort()
		return
	}
	fmt.Println("请求的token:", token)

	// 执行函数
	c.Next()
	// 更新token超时时效

	// 中间件执行完后续的一些事情
	status := c.Writer.Status()
	t2 := time.Since(t)
	fmt.Printf("认证请求中间件执行完毕: %v ,time: %v\n", status, t2)
}
