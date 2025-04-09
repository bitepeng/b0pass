package app

import (
	"b0go/core/engine"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取token/login
// 权限code: ro, rw
func Login(c *gin.Context) {
	code := c.Query("code")
	auth := ""
	if code == "" {
		engine.ERR("code is empty", c)
		return
	}
	if code != config.CodeReadOnly && code != config.CodeReadWrite {
		engine.ERR("code is invalid", c)
		return
	} else {
		// login success
		if code == config.CodeReadOnly {
			auth = "ro"
			code = "ro:" + code
			engine.Println("code is read only")
		}
		if code == config.CodeReadWrite {
			auth = "rw"
			code = "rw:" + code
			engine.Println("code is read write")
		}
	}
	token, err := engine.GenerateToken(code)
	if err != nil {
		engine.ERR(err.Error(), c)
		return
	}
	engine.OK("OK", auth+":"+token, c)
}

// JWTMiddleware 基于JWT的认证中间件
func JWTMiddleware(mode string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 如果未设置任何验证码，直接放行
		if config.CodeReadOnly == "" && config.CodeReadWrite == "" {
			c.Next()
			return
		} else {
			// 客户端携带Token方式 1.请求头 2.请求体 3.URI
			// Token放在Header的token中
			authHeader := c.Request.Header.Get("token")
			if authHeader == "" {
				authHeader = c.Query("token")
			}
			if authHeader == "" {
				c.JSON(http.StatusOK, gin.H{"code": 401, "msg": "请求缺少token信息"})
				c.Abort()
				return
			}
			// 检查Token
			mc, err := engine.ParseToken(authHeader)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 401, "msg": "请求的token信息无效"})
				c.Abort()
				return
			}

			// 权限检查
			user := mc.User
			if mode == "ro" && !(strings.HasPrefix(user, "ro:") || strings.HasPrefix(user, "rw:")) {
				c.JSON(http.StatusOK, gin.H{"code": 403, "msg": "只读权限不足"})
				c.Abort()
				return
			} else if mode == "rw" && !strings.HasPrefix(user, "rw:") {
				c.JSON(http.StatusOK, gin.H{"code": 403, "msg": "读写权限不足"})
				c.Abort()
				return
			}

			// 将当前user保存到请求的上下文c
			// 用c.Get("user")获取当前请求用户信息
			c.Set("user", mc.User)
			c.Next()
		}
	}
}

// 释放dll文件
func putDll() {
	if runtime.GOOS == "windows" && !config.Live {
		_, errNow := os.ReadFile("zlib1.dll")
		if errNow != nil {
			dll, err := uiFS.ReadFile("ui/dist/dll/zlib1.dll")
			if err != nil {
				log.Println("zlib1.dll err:", err)
			} else {
				os.WriteFile("zlib1.dll", dll, 0777)
			}
		}
	}
}

// 干净安全的路径
func cleanPath(p string) string {
	p = path.Clean(p)
	p = strings.ReplaceAll(p, "\\", "/")
	p = strings.ReplaceAll(p, "..", "")
	p = strings.ReplaceAll(p, "/./", "/")
	p = strings.ReplaceAll(p, "//", "/")
	return p
}

// 干净连接路径
func cleanPathJoin(p, f string) string {
	p = cleanPath(p)
	f = cleanPath(f)
	if strings.HasSuffix(p, "/") {
		return cleanPath(p + f)
	}
	return cleanPath(p + "/" + f)
}

// 检查路径是否在锁定上传目录内
func checkLockUploadDir(f string, c *gin.Context) bool {
	RootPath := cleanPath(config.Path)
	lockUploadDir := cleanPathJoin(RootPath, config.LockUploadDir)
	if config.LockUploadDir != "" {
		if !strings.HasPrefix(f, lockUploadDir) {
			engine.ERR("该目录不能操作", c)
			return false
		}
	}
	return true
}
