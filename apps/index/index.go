package index

import (
	"b0pass/boot"
	conf2 "b0pass/library/conf"
	"b0pass/library/fileinfos"
	"b0pass/library/ipaddress"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/os/gfile"
	"strconv"
	"time"
)

var (
	filePath string
)

type Controller struct {
	gmvc.Controller
}

func (c *Controller) Index() {
	c.View.Assign("times", time.Now().Unix())
	_ = c.View.Display("index.html")
}

func (c *Controller) FileLists() {
	// Ip lists
	port := boot.ServPort
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+strconv.Itoa(port))
	}
	c.View.Assign("ips", ips)
	// path
	//	pathRoot := fileinfos.GetRootPath() + "/files/"
	//pathRoot := "D:/files/"
	if gfile.Exists(fileinfos.GetRootPath() + "/config.conf") {
		conf := conf2.InitConfig(fileinfos.GetRootPath() + "/config.conf")
		filePath = conf["filePath"]
	} else {
		filePath = fileinfos.GetRootPath() + "/files/"
	}
	c.View.Assign("path_root", filePath)
	// file lists
	fprPath := c.Request.GetString("path")
	var fpPath string
	if fprPath != "" {
		fpPath = fprPath + "/*"
	} else {
		fpPath = "/*"
	}
	fp := filePath + fpPath
	flists := fileinfos.ListDirData(fp, fprPath)
	c.View.Assign("flists", flists)
	// views
	_ = c.View.Display("file-lists.html")
}
