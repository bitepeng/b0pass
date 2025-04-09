package app

import (
	"b0go/apps/pass/lib/chat"
	"b0go/core/engine"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// 注册ws路由
func routeWs() {
	hub := chat.NewHub()
	go hub.Run()
	engine.Gin.GET("/ws", func(c *gin.Context) {
		chat.ServeWs(hub, c)
	})
}

// 注册应用路由
func routeApi() {
	//无需鉴权
	GET("/login", "{code=权限code}", "login", Login)
	GET("/ping", "{}", "连通性测试")

	//需要鉴权ro
	GETX("/read-config", "{}", "读取配置(ro)", ReadConfig)
	GETX("/read-ip", "{}", "读取IP(ro)", ReadIP)
	GETX("/node-tree", "{}", "目录树结构(ro)", NodeTree)
	GETX("/file-count", "{f=相对路径}", "文件数量(ro)", FileCount)
	GETX("/file-list", "{f=相对路径,[t=需要的类型]}", "文件列表(ro)", FileList)
	GETX("/file-content", "{f=相对路径,文件名称}", "文件内容(ro)", FileContent)
	GETX("/file-download", "{f=相对路径}", "文件列表(ro)", FileDownload)

	//需要鉴权rw
	POSTX("/cmd-open", "{}", "命令行打开(rw)", CmdOpen)
	POSTX("/cmd-key", "{}", "主电脑键盘(rw)", CmdKey)
	POSTX("/node-add", "{f=相对路径(结尾带“/”为创建目录,否则为创建文件)}", "添加目录(rw)", NodeAdd)
	POSTX("/node-rename", "{f=原路径,n=新路径}", "重命名节点(rw)", NodeRename)
	POSTX("/node-delete", "{f=相对路径}", "删除节点(rw)", NodeRemove)
	POSTX("/file-upload", "{post file}", "大文件上传(rw)", FileUpload)
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
	engine.Router(appId, "GET", url, param, "(Auth)"+title, JWTMiddleware("ro"), handle)
}

// POSTX
func POSTX(url, param, title string, handle gin.HandlerFunc) {
	engine.Router(appId, "POST", url, param, "(Auth)"+title, JWTMiddleware("rw"), handle)
}
