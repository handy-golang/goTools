package mGin

import (
	"net/http"

	"github.com/EasyGolang/goTools/mRes"
	"github.com/gin-gonic/gin"
)

func CrossDomain(c *gin.Context) {
	// 请求方法
	method := c.Request.Method
	// 指明哪些请求源被允许访问资源，值可以为 "*"（允许访问所有域），"null"，或者单个源地址。
	c.Header("Access-Control-Allow-Origin", "*")
	// 对于预请求来说，哪些请求方式可以用于实际的请求。
	c.Header("Access-Control-Allow-Methods", "*")
	// 对于预请求来说，指明哪些头信息可以安全的暴露给 CORS API 规范的 API
	c.Header("Access-Control-Expose-Headers", "*")
	// 允许Header
	c.Header("Access-Control-Allow-Headers", "*")
	// 允许客户端传递校验信息比如 cookie (重要)
	c.Header("Access-Control-Allow-Credentials", "true")

	// 设置返回格式是json
	c.Set("Content-Type", "application/json; charset=utf-8")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.JSON(http.StatusOK, mRes.Response(0, "OPTIONS").WithData("OPTIONS"))
		c.Abort()
	}
}
