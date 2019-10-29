package boot

import (
	"b0pass/library/fileinfos"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gres"
	"time"
)

var (
	PathRoot string
)

// 用于应用初始化。
func init() {

	go func() {

		// 根目录
		PathRoot = fileinfos.GetRootPath()
		fmt.Println("ROOT:", PathRoot)
		gres.Dump()

		v := g.View()
		c := g.Config()
		s := g.Server()

		// 模板引擎配置
		_ = v.AddPath("template")
		v.SetDelimiters("${", "}")

		// glog配置
		logpath := c.GetString("setting.logpath")
		_ = glog.SetPath(logpath)
		//glog.SetStdoutPrint(true)

		// Web Server配置
		s.SetIndexFolder(true)
		s.SetServerRoot("public")
		s.SetLogPath(logpath)
		s.SetReadTimeout(3 * 60 * time.Second)
		s.SetWriteTimeout(3 * 60 * time.Second)
		s.SetIdleTimeout(3 * 60 * time.Second)
		s.SetMaxHeaderBytes(1024 * 1024 * 200)
		s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)
		//s.SetErrorLogEnabled(true)
		//s.SetAccessLogEnabled(true)
		s.SetPort(c.GetInt("setting.port"))
		s.SetDumpRouteMap(false)

		// 文件根目录
		filePath := PathRoot + "/files"
		if !gfile.Exists(filePath) {
			if err := gfile.Mkdir(filePath); err != nil {
				panic(err)
			}
		}
		s.AddStaticPath("/files", filePath)

		// Run Server
		g.Server().Run()
	}()

}
