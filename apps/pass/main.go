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
	Live          bool
	Path          string //文件根目录路径
	CodeReadOnly  string //只读权限code
	CodeReadWrite string //读写权限code
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
	routeWs()
	putDll()
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
	//engine.Gin.StaticFS("/files", http.Dir(config.Path))
	// files静态路由添加JWT校验
	filesGroup := engine.Gin.Group("/files")
	filesGroup.Use(JWTMiddleware("ro")) // 使用只读权限校验
	filesGroup.StaticFS("/", http.Dir(config.Path))
}

// 注册应用路由
func routeApi() {
	//鉴权token
	GET("/login", "{code=权限code}", "login", Login)
	GETX("/ping", "{}", "连通性测试", Ping, "ro")
	GETX("/read-config", "{}", "读取配置(ro)", ReadConfig, "ro")
	GETX("/read-ip", "{}", "读取IP(ro)", ReadIP, "ro")
	GETX("/cmd-open", "{}", "命令行打开(rw)", CmdOpen, "rw")
	GETX("/cmd-key", "{}", "主电脑键盘(rw)", CmdKey, "rw")

	GETX("/node-tree", "{}", "目录树结构(ro)", NodeTree, "ro")
	GETX("/node-add", "{f=相对路径(结尾带“/”为创建目录,否则为创建文件)}", "添加目录(rw)", NodeAdd, "rw")
	GETX("/node-rename", "{f=原路径,n=新路径}", "重命名节点(rw)", NodeRename, "rw")
	GETX("/node-delete", "{f=相对路径}", "删除节点(rw)", NodeRemove, "rw")

	GETX("/file-count", "{f=相对路径}", "文件数量(ro)", FileCount, "ro")
	GETX("/file-list", "{f=相对路径,[t=需要的类型]}", "文件列表(ro)", FileList, "ro")
	GETX("/file-content", "{f=相对路径,文件名称}", "文件内容(ro)", FileContent, "ro")
	GETX("/file-download", "{f=相对路径}", "文件列表(ro)", FileDownload, "ro")

	POSTX("/file-upload", "{post file}", "大文件上传(rw)", FileUpload, "rw")
}
