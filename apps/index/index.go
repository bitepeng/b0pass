package index

import (
	"b0pass/boot"
	"b0pass/library/fileinfos"
	"b0pass/library/ipaddress"
	"github.com/gogf/gf/frame/gmvc"
	"strconv"
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
	port := boot.ServPort
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+strconv.Itoa(port))
	}
	c.View.Assign("ips",ips)
	// file lists
	fp := fileinfos.GetRootPath() + "/files/*"
	flists := fileinfos.ListDirData(fp)
	c.View.Assign("flists",flists)
	// views
	_ = c.View.Display("file-lists.html")
}