package router

import (
	"b0pass/apps/api"
	"b0pass/apps/index"
	"b0pass/apps/sync"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// Index
	s.BindController("/", new(index.Controller))

	// Chat
	//s.BindController("/chat", new(chat.Controller))
	s.BindController("/sync", new(sync.Controller))

	// Api
	s.Group("/api", func(g *ghttp.RouterGroup) {
		//cors
		g.Middleware(MiddlewareCORS)
		//file
		g.POST("/upload", api.Upload)
		g.GET("/lists", api.Lists)
		g.GET("/delete", api.Delete)
		g.GET("/dump", api.Dump)
		g.GET("/upload", api.UploadShow)
		//server
		g.GET("/sip", api.GetIp)
		g.ALL("/subpath", api.GetSubPath)
		g.ALL("/textdata", api.GetTextData)
		g.GET("/openurl",api.OpenUrl)
	})

}
