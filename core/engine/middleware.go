package engine

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

/****** Cors MiddleWare *******/

// 处理跨域请求,支持options访问
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

/****** Token MiddleWare *******/

var (
	TokenExpire = time.Hour * 24
	TokenSecret = []byte("00xda0d8f6x9n1x8")
)

type MyClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return TokenSecret, nil
		})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTMiddleware 基于JWT的认证中间件
func JWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token方式 1.请求头 2.请求体 3.URI
		// Token放在Header的xtoken中
		authHeader := c.Request.Header.Get("token")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{"code": 401, "msg": "请求缺少token信息"})
			c.Abort()
			return
		}
		// 检查Token
		mc, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 401, "msg": "请求的token信息无效"})
			c.Abort()
			return
		}
		// 将当前user保存到请求的上下文c
		// 用c.Get("user")获取当前请求用户信息
		c.Set("user", mc.User)
		c.Next()
	}
}
