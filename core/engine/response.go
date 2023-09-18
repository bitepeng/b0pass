package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OKERR 成功失败自适应返回 code=0
func OKERR(msg, errmsg string, data interface{}, c *gin.Context) {
	code := 0
	if errmsg != "" {
		code = 400
		msg = errmsg
	}
	JSON(code, msg, data, c)
}

// OK 成功返回 code=0
func OK(msg string, data interface{}, c *gin.Context) {
	JSON(0, msg, data, c)
}

// ERROR 失败返回 code=400
func ERR(msg string, c *gin.Context) {
	JSON(400, msg, nil, c)
}

// ECHO 直接返回
func ECHO(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

// JSON 通用返回参数
// code=0 成功OK
// code=400 错误ERROR
// code=401 需要登录
func JSON(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// PAGE 返回分页数据
func PAGE(count int64, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
	})
}
