package api

import (
	"b0pass/boot"
	"b0pass/library/ipaddress"
	nustdbs "b0pass/library/nutsdbs"
	"b0pass/library/openurl"
	"b0pass/library/response"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

// GetIp 获取IP地址
func GetIp(r *ghttp.Request) {
	port := boot.ServPort
	ip, _ := ipaddress.GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp+":"+strconv.Itoa(port))
	}
	response.JSON(r, 0, "ok", ips)
}

// GetPathSub 上传目录记忆功能
func GetSubPath(r *ghttp.Request){
	getData:=r.GetString("path")
	dbKey:="files_path"
	if getData!=""{
		nustdbs.DBs.SetData(dbKey,getData)
	}
	dbData :=nustdbs.DBs.GetData(dbKey)
	/*if filesPath =="" {
		dbData:=time.Now().Format("2006-01-02")
		nustdbs.DBs.SetData("files_path",dbData)
	}*/
	response.JSON(r, 0, "ok", dbData)
}

// GetTextData 文本内容共享
func GetTextData(r *ghttp.Request){
	getData:=r.GetString("data")
	dbKey:="data_text"
	if getData!=""{
		nustdbs.DBs.SetData(dbKey,getData)
	}
	dbData :=nustdbs.DBs.GetData(dbKey)
	response.JSON(r, 0, "ok", dbData)
}

// OpenUrl 打开本地url
func OpenUrl(r *ghttp.Request){
	getUrl:=r.GetString("url")
	_ = openurl.Open(getUrl)
}