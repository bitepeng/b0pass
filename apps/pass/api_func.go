package app

import (
	"b0go/apps/pass/lib/keys"
	"b0go/core/engine"
	"b0go/core/tools/cmd"
	"b0go/core/tools/nets"
	"log"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ping 主电脑连通性测试
func Ping(c *gin.Context) {
	engine.OK("OK", "pong", c)
}

// ReadConfig 读取配置
func ReadConfig(c *gin.Context) {
	config_ := config
	config_.CodeReadOnly = ""
	config_.CodeReadWrite = ""
	engine.OK("OK", config_, c)
}

// ReadIP 读取配置
func ReadIP(c *gin.Context) {
	ip := nets.GetOutBoundIP()
	engine.OK("OK", ip, c)
}

// CmdOpen 使用命令行打开
func CmdOpen(c *gin.Context) {
	RootPath := strings.ReplaceAll(config.Path, "\\", "/")
	f := c.Query("f")
	f = strings.ReplaceAll(f, "\\", "/")
	ext := strings.ToUpper(path.Ext(f))
	engine.Println(ext)
	if ext == ".BAT" || ext == ".CMD" || ext == ".EXE" {
		engine.ERR("该文件暂不支持打开", c)
	} else {
		cmd.Open(RootPath + f)
	}
}

// CmdKey 主电脑键盘
func CmdKey(c *gin.Context) {
	k := c.Query("k")
	log.Printf("::CmdKey:: %s", k)
	keys.SendKey(k)
	engine.OK("OK", nil, c)
}
