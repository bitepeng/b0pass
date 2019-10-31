package router

import (
	"b0pass/apps/api"
	"b0pass/apps/index"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// Index
	s.BindController("/", new(index.Controller))

	// Api
	s.Group("/api", func(g *ghttp.RouterGroup) {
		//跨域设置
		g.Middleware(MiddlewareCORS)
		//文件上传
		g.POST("/uploado", api.Upload)
		g.POST("/upload", api.Uploadx)
		g.GET("/lists", api.Lists)
		g.GET("/delete", api.Delete)
		g.GET("/sip", api.GetIp)
		g.GET("/dump", api.Dump)
		g.GET("/upload", api.UploadShow)
	})
}
