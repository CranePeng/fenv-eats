package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 全局，以后迁移到fenv-middleware项目里面
func GlobalMidWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 请求设置独有的trace id，用来数据归纳

		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()

		// 获取验证码中间件 session赋予的值
		value, exitBool := c.Get("github.com/gin-contrib/sessions")
		if exitBool {
			fmt.Println("session中间件赋予数据：", value)
		}

		fmt.Println("全局中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 跨域中间件，可以参考https://www.cnblogs.com/you-men/p/14054348.html
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
