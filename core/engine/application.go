package engine

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

// App 所有的应用配置
var App = make(map[string]*AppConfig)

// APP Type类型定义
const (
	APP_NONE       = 0      //独立应用
	APP_SUBSCRIBER = 1      //订阅者应用
	APP_PUBLISHER  = 1 << 1 //发布者应用
	APP_HOOK       = 1 << 2 //钩子应用
	APP_APP        = 1 << 3 //应用应用
)

// AppConfig 应用配置定义
type AppConfig struct {
	Name    string                         //应用名称
	Type    byte                           //类型
	Config  interface{}                    //应用配置
	UIDir   string                         //界面目录
	UIFS    fs.FS                          //界面目录FS
	Dir     string                         //应用代码路径
	Run     func()                         //应用启动函数
	Version string                         //应用版本
	Router  map[string][]map[string]string //应用路由表
}

// AppInstall 安装应用
func AppInstall(opt *AppConfig) {
	App[opt.Name] = opt
	opt.Router = make(map[string][]map[string]string)
	//Dir
	_, appFilePath, _, _ := runtime.Caller(1)
	opt.Dir = filepath.Dir(appFilePath)
	//UIDir
	ui := filepath.Join(opt.Dir, "ui", "dist")
	if _, err := os.Stat(ui); err == nil || os.IsExist(err) {
		if opt.UIDir == "" {
			opt.UIDir = ui
		}
	}
	if parts := strings.Split(opt.Dir, "@"); len(parts) > 1 {
		opt.Version = parts[len(parts)-1]
	}
	//Print(aurora.Green("install app"), aurora.BrightCyan(opt.Name), aurora.BrightBlue(opt.Version))
}

// GET 路由处理
func GET(appId, url, param, title string, handle ...gin.HandlerFunc) {
	Router(appId, "GET", url, param, title, handle...)
}

// POST 路由处理
func POST(appId, url, param, title string, handle ...gin.HandlerFunc) {
	Router(appId, "POST", url, param, title, handle...)
}

// Router 路由处理
func Router(appId, method, url, param, title string, handle ...gin.HandlerFunc) {
	//log.Println("【router】", appId, url, title, method)
	opt := App[appId]
	url = "/" + appId + url
	urlmap := make(map[string]string)
	urlmap["url"] = url
	urlmap["param"] = param
	urlmap["title"] = title
	opt.Router[method] = append(opt.Router[method], urlmap)
	switch method {
	case "GET":
		Gin.GET(url, handle...)
	case "POST":
		Gin.POST(url, handle...)
	case "DELETE":
		Gin.DELETE(url, handle...)
	case "PUT":
		Gin.POST(url, handle...)
	default:
		Gin.Any(url, handle...)
	}
}
