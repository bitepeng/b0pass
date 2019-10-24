package api

import (
	"b0pass/library/ipaddress"
	"b0pass/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GetIp(r *ghttp.Request) {
	c := g.Config()
	port := c.GetString("setting.port")
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+port)
	}
	response.JSON(r, 0, "ok", ips)
}
