package app

import (
	"b0go/core/engine"
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
)

// APP:AppConfig
type AppConfig struct {
	Live bool
	Path string //文件根目录路径
}

// APP:VAR
var (
	app    *engine.AppConfig
	appId  = "pass"
	config = new(AppConfig)

	uiPath string
	//go:embed ui/dist
	uiFS embed.FS
)

// APP:INIT
func init() {
	uiDist, _ := fs.Sub(uiFS, "ui/dist")
	app = &engine.AppConfig{
		Name:   appId,
		Type:   engine.APP_APP,
		Config: config,
		UIFS:   uiDist,
		Run:    run,
	}
	engine.AppInstall(app)
	uiPath = filepath.Join(app.Dir, "ui", "dist")
}

// APP:RUN
func run() {
	engine.Print(aurora.Green("App pass loaded"), aurora.BrightCyan(config))
	engine.Gin.Use(engine.CorsMiddleware())
	routeStatic(config.Live)
	routeApi()
}

// 注册静态路由
func routeStatic(live bool) {
	if live {
		engine.Gin.Static("/index", uiPath)
	} else {
		uiDist, _ := fs.Sub(uiFS, "ui/dist")
		engine.Gin.StaticFS("/index", http.FS(uiDist))
	}
	//默认首页
	engine.Gin.GET("/", func(c *gin.Context) {
		//c.Redirect(http.StatusMovedPermanently, "/index")
		c.Writer.Write([]byte("<script>location.href='/app/pass/'</script>"))
	})
	//files静态
	engine.Gin.StaticFS("/files", http.Dir(config.Path))
}

// 注册应用路由
func routeApi() {
	GETX("/ping", "{}", "连通性测试", Ping)
	engine.GET(appId, "/read-config", "{}", "读取配置", ReadConfig)
	engine.GET(appId, "/read-ip", "{}", "读取IP", ReadIP)
	engine.GET(appId, "/cmd-open", "{}", "命令行打开", CmdOpen)
	engine.GET(appId, "/cmd-key", "{}", "主电脑键盘", CmdKey)

	engine.GET(appId, "/node-tree", "{}", "目录树结构", NodeTree)
	engine.GET(appId, "/node-add", "{f=相对路径(结尾带“/”为创建目录,否则为创建文件)}", "添加目录", NodeAdd)
	engine.GET(appId, "/node-rename", "{f=原路径,n=新路径}", "重命名节点", NodeRename)
	engine.GET(appId, "/node-delete", "{f=相对路径}", "删除节点", NodeRemove)

	engine.GET(appId, "/file-count", "{f=相对路径}", "文件数量", FileCount)
	engine.GET(appId, "/file-list", "{f=相对路径,[t=需要的类型]}", "文件列表", FileList)
	engine.GET(appId, "/file-content", "{f=相对路径,文件名称}", "文件内容", FileContent)
	engine.GET(appId, "/file-download", "{f=相对路径}", "文件列表", FileDownload)

	engine.POST(appId, "/file-upload", "{post file}", "大文件上传", FileUpload)
}

/***** HttpRequest *****/

// GET
func GET(url, param, title string, handle ...gin.HandlerFunc) {
	engine.Router(appId, "GET", url, param, title, handle...)
}

// POST
func POST(url, param, title string, handle ...gin.HandlerFunc) {
	engine.Router(appId, "POST", url, param, title, handle...)
}

// GETX
func GETX(url, param, title string, handle gin.HandlerFunc) {
	engine.Router(appId, "GET", url, param, "(Auth)"+title, engine.JWTMiddleware(), handle)
}

// POSTX
func POSTX(url, param, title string, handle gin.HandlerFunc) {
	engine.Router(appId, "POST", url, param, "(Auth)"+title, engine.JWTMiddleware(), handle)
}
