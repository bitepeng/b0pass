package boot

import (
	"b0pass/library/fileinfos"
	"b0pass/library/nutsdbs"
	"flag"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"time"
)

var (
	PathRoot string
	ServPort int
)

func ExecArgs(){
	flag.Parse()
	if ServPort<=0{
		ServPort=g.Config().GetInt("setting.port")
	}
}

// 用于应用初始化。
func init() {

	// 分析CLI参数
	flag.IntVar(&ServPort,"p",8899,"-p for Server Port(default=8899)")
	ExecArgs()

	// 资源根目录
	PathRoot = fileinfos.GetRootPath()

	go func() {

		// APP核心引擎
		v := g.View()
		c := g.Config()
		s := g.Server()

		// 加载动作缓冲
		time.Sleep(3000 * time.Millisecond)

		// 模板引擎配置
		_ = v.AddPath("template")
		v.SetDelimiters("${", "}")

		// glog配置
		logpath := c.GetString("setting.logpath")
		_ = glog.SetPath(logpath)
		glog.SetStdoutPrint(true)

		// Web Server配置
		s.SetIndexFolder(true)
		s.SetServerRoot("public")
		s.SetLogPath(logpath)
		s.SetReadTimeout(3 * 60 * time.Second)
		s.SetWriteTimeout(3 * 60 * time.Second)
		s.SetIdleTimeout(3 * 60 * time.Second)
		s.SetMaxHeaderBytes(32*1024)
		s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)
		s.SetErrorLogEnabled(true)
		s.SetAccessLogEnabled(true)
		s.SetPort(ServPort)
		s.SetDumpRouteMap(false)

		// 文件根目录
		filePath := PathRoot + "/files"
		if !gfile.Exists(filePath) {
			if err := gfile.Mkdir(filePath); err != nil {
				panic(err)
			}
		}
		s.AddStaticPath("/files", filePath)

		// 内存数据库
		defer func() {
			nustdbs.DBs.CloseDB()
		}()
		nustdbs.DBs.SetData("files_root",filePath)

		// Run Server
		g.Server().Run()
	}()

}
