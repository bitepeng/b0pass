package gateway

import (
	"b0go/core/engine"
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
)

type AppConfig struct {
	ListenAddr string
	Domain     string
	Live       bool
}

var (
	appId  = "gateway"
	config = new(AppConfig)

	docsPath string
	//go:embed doc/dist
	docsPathFS embed.FS

	//go:embed ui/dist
	uiFS embed.FS
)

func init() {
	uiDist, _ := fs.Sub(uiFS, "ui/dist")
	app := &engine.AppConfig{
		Name:   appId,
		Type:   engine.APP_HOOK,
		Config: config,
		UIFS:   uiDist,
		Run:    run,
	}
	engine.AppInstall(app)
	docsPath = filepath.Join(app.Dir, "doc", "dist")
}

/*
---------------------------------------------------
- /						# 首页 /docs/index.md
- /docs/%mdfile.html%	# md文档 /docs/%mdfile.md%
---------------------------------------------------
- /dev/doc				# api doc 文档页面
- /dev/api				# api json
---------------------------------------------------
- /app/%appid% 			# 具体应用首页
- /app/%appid%/%act% 	# 具体应用api json
---------------------------------------------------
*/
func run() {
	//注册应用ui路由
	engine.Print(aurora.Green("App gateway loaded"), aurora.BrightCyan(config))
	addAppStaticRoute(config.Live)

	//动态路由挂载
	engine.GET(appId, "/config", "{}", "read config", ReadConfig)
	engine.Gin.GET("/dev/api", getApp)

	//启动httpd服务
	engine.Addr = config.ListenAddr
	engine.Domain = config.Domain
	engine.Gin.Run(config.ListenAddr)
}

// 读取配置信息
func ReadConfig(c *gin.Context) {
	engine.OK("OK", config, c)
}

// 注册模块静态文件路由 /app/moduleName
func addAppStaticRoute(live bool) {
	//默认首页
	/*engine.Gin.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/docs/index.html"
		engine.Gin.HandleContext(c)
	})*/
	if live {
		//docs
		engine.Gin.Static("/dev/doc", docsPath)
		//app/appid
		for name, config := range engine.App {
			engine.Gin.Static(fmt.Sprintf("/app/%s", name), config.UIDir)
		}
	} else {
		//docs
		docDist, _ := fs.Sub(docsPathFS, "doc/dist")
		engine.Gin.StaticFS("/dev/doc", http.FS(docDist))
		//app/appid
		for name, config := range engine.App {
			engine.Gin.StaticFS(fmt.Sprintf("/app/%s", name), http.FS(config.UIFS))
		}
	}
}

// AppInfo 应用信息JSON结构
type AppInfo struct {
	Name    string                         //应用名称
	Type    byte                           //类型
	Config  string                         //应用配置
	UIDir   string                         //界面路径
	Dir     string                         //应用代码目录
	ReadMe  string                         //README.md
	Version string                         //应用版本
	Router  map[string][]map[string]string //应用路由表
}

// getApp 获取App实时信息
func getApp(c *gin.Context) {
	apps := make(map[string]*AppInfo)
	for _, app := range engine.App {
		p := &AppInfo{
			app.Name,
			app.Type,
			"",
			app.UIDir,
			app.Dir,
			"",
			app.Version,
			app.Router,
		}
		if bytes, err := ioutil.ReadFile(filepath.Join(app.Dir, "README.md")); err == nil {
			p.ReadMe = string(bytes[0:100])
		}
		var out bytes.Buffer
		if toml.NewEncoder(&out).Encode(app.Config) == nil {
			p.Config = out.String()
		}
		apps[app.Name] = p
	}
	c.JSON(200, apps)
}
