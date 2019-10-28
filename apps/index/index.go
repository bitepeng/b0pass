package index

import (
	"b0pass/library/fileinfos"
	"b0pass/library/ipaddress"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"
)

type Controller struct {
	gmvc.Controller
}

func (c *Controller) Index() {
	c.View.Assign("times",time.Now().Unix())
	_ = c.View.Display("index.html")
}

func (c *Controller) FileLists() {
	// Ip lists
	port := g.Config().GetString("setting.port")
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+port)
	}
	c.View.Assign("ips",ips)
	// file lists
	fp := fileinfos.GetRootPath() + "/files/*"
	flists := fileinfos.ListDirData(fp)
	c.View.Assign("flists",flists)
	// views
	_ = c.View.Display("file-lists.html")
}