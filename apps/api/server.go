package api

import (
	"b0pass/boot"
	"b0pass/library/fileinfos"
	"b0pass/library/ipaddress"
	"b0pass/library/openurl"
	"b0pass/library/response"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

// OpenUrl 打开本地url
func OpenUrl(r *ghttp.Request){
	getUrl:=r.GetString("url")
	_ = openurl.Open(getUrl)
}

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
	saveData(r,"path","data_path")
}

// GetTextData 文本内容共享
func GetTextData(r *ghttp.Request){
	saveData(r,"data","data_text")
}

// saveData 保存数据
func saveData(r *ghttp.Request,getkey string,dbkey string){
	getData:=r.GetPostString(getkey)
	getCode:=r.GetPostString("code")
	if getCode=="1"{
		fileinfos.Set(dbkey,getData)
	}
	dbData :=fileinfos.Get(dbkey)
	//dbData:=time.Now().Format("2006-01-02")
	response.JSON(r, 0, "ok", dbData)
}