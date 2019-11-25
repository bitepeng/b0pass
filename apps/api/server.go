package api

import (
	"b0pass/boot"
	"b0pass/library/fileinfos"
	"b0pass/library/ipaddress"
	nustdbs "b0pass/library/nutsdbs"
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

func GetPath(r *ghttp.Request) {
	path:=fileinfos.GetRootPath()
	response.JSON(r, 0, "ok", path)
}

func GetPathSub(r *ghttp.Request){
	getPath:=r.GetString("path")
	filesPath :=nustdbs.DBs.GetData("files_path")
	if getPath!="" {
		filesPath=getPath
		nustdbs.DBs.SetData("files_path",filesPath)
	}
	/*if filesPath =="" {
		dates:=time.Now().Format("2006-01-02")
		nustdbs.DBs.SetData("files_path",dates)
	}*/
	response.JSON(r, 0, "ok", filesPath)
}
