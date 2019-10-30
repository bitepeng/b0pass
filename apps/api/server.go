package api

import (
	"b0pass/boot"
	"b0pass/library/ipaddress"
	"b0pass/library/response"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

func GetIp(r *ghttp.Request) {
	port := boot.ServPort
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+strconv.Itoa(port))
	}
	response.JSON(r, 0, "ok", ips)
}
