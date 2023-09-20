package app

import (
	"b0go/apps/pass/lib/files"
	"b0go/apps/pass/lib/stream"
	"b0go/core/engine"
	"b0go/core/tools/cmd"
	"b0go/core/tools/nets"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ping 主电脑连通性测试
func Ping(c *gin.Context) {
	engine.OK("OK", true, c)
}

// ReadConfig 读取配置
func ReadConfig(c *gin.Context) {
	engine.OK("OK", config, c)
}

// ReadIP 读取配置
func ReadIP(c *gin.Context) {
	ip := nets.GetOutBoundIP()
	engine.OK("OK", ip, c)
}

// CmdOpen 使用命令行打开
func CmdOpen(c *gin.Context) {
	RootPath := config.Path
	f := c.Query("f")
	ext := strings.ToUpper(path.Ext(f))
	engine.Println(ext)
	if ext == ".BAT" || ext == ".CMD" || ext == ".EXE" {
		engine.ERR("该文件暂不支持打开", c)
	} else {
		cmd.Open(RootPath + f)
	}
}

// NodeTree 目录树结构
func NodeTree(c *gin.Context) {
	f := c.Query("f")
	listMap := files.NodeTree(config.Path, f)
	engine.OK("OK", listMap, c)
}

// NodeRename 重命名目录
func NodeRename(c *gin.Context) {
	RootPath := config.Path
	f := c.Query("f")
	n := c.Query("n")
	if f == "" || n == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	err := files.NodeRename(RootPath+f, RootPath+n)
	if err != nil {
		engine.ERR(err.Error(), c)
		return
	}
	engine.OK("OK", nil, c)
}

// NodeRemove 删除目录
func NodeRemove(c *gin.Context) {
	RootPath := config.Path
	f := c.Query("f")
	if f == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	err := files.NodeRemove(RootPath + f)
	if err != nil {
		engine.ERR(err.Error(), c)
		return
	}
	engine.OK("OK", nil, c)
}

// NodeAdd 添加目录
func NodeAdd(c *gin.Context) {
	RootPath := config.Path
	f := c.Query("f")
	if f == "" {
		engine.ERR("路径不能为空:(f = 结尾带“/”为创建目录,否则为创建文件)", c)
		return
	}
	filePath := RootPath + f
	filePath = strings.ReplaceAll(filePath, "//", "/")
	log.Println("::NodeAdd::", filePath)
	err := files.NodeAdd(filePath)
	if err != nil {
		engine.ERR(err.Error(), c)
	}
	engine.OK("OK", nil, c)
}

func FileCount(c *gin.Context) {
	RootPath := config.Path
	listMap := files.GetCounts(RootPath)
	engine.OK("OK", listMap, c)
}

// FileList 文件列表
func FileList(c *gin.Context) {
	RootPath := config.Path
	f := c.Query("f")
	t := c.DefaultQuery("t", "")
	listMap := files.GetDirTree(RootPath, RootPath+f, "", t)
	engine.OK("OK", listMap, c)
}

// FileContent 文件内容
func FileContent(c *gin.Context) {
	RootPath := config.Path
	f := c.Query("f") //!strings.HasPrefix(f, RootPath)
	if f == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	RootPath = path.Clean(RootPath + f)
	data, err := files.GetData(RootPath)
	if err != nil {
		engine.ERR(err.Error(), c)
		return
	}
	engine.OK("OK", string(data), c)
}

// FileDownload 文件下载
func FileDownload(c *gin.Context) {
	RootPath := config.Path
	f := c.Query("f")
	if f == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	filePath := RootPath + f
	//获取文件的名称
	fileName := path.Base(filePath)

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	//fmt.Println(filePath)
	c.File(filePath)
}

// FileUpload 上传文件
func FileUpload(c *gin.Context) {
	lens, _ := strconv.Atoi(c.Request.Header["Content-Length"][0])
	log.Println("FileUpload::::", lens)
	if lens > 4096 {
		FileUploadBig(c)
	} else {
		FileUploadTiny(c)
	}
}

// FileUploadTiny 小文件上传
func FileUploadTiny(c *gin.Context) {
	RootPath := config.Path
	f := c.DefaultQuery("f", "/")
	RootPath = RootPath + f
	files.NodeAdd(RootPath)
	log.Println("FileUploadBig:::", RootPath)
	//上传文件
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, RootPath+file.Filename)
	//上传成功
	engine.OK("上传成功", "", c)
}

// FileUploadBig 大文件上传
func FileUploadBig(c *gin.Context) {

	RootPath := config.Path
	f := c.DefaultQuery("f", "/")
	RootPath = RootPath + f
	files.NodeAdd(RootPath)
	log.Println("FileUploadBig:::", RootPath)

	content_type_, has_key := c.Request.Header["Content-Type"]
	if !has_key {
		engine.ERR("请先选择文件", c)
		return
	}
	content_type := content_type_[0]
	const BOUNDARY string = "; boundary="
	loc := strings.Index(content_type, BOUNDARY)
	if len(content_type_) != 1 || loc == -1 {
		engine.ERR("请先选择文件", c)
		return
	}
	boundary := []byte(content_type[(loc + len(BOUNDARY)):])
	log.Printf("file boundary: [%s]\n", boundary)
	read_data := make([]byte, 1024*12)
	var read_total int = 0
	for {
		//解析文件头信息
		file_header, file_data, err := stream.ParseFromHead(
			read_data, read_total,
			append(boundary, []byte("\r\n")...), c.Request.Body,
		)
		if err != nil {
			engine.ERR("parse from fail: "+err.Error(), c)
			return
		}
		//创建保存文件
		log.Printf("save file: [%s]\n", RootPath+file_header.FileName)
		f, err := os.Create(RootPath + file_header.FileName)
		if err != nil {
			engine.ERR("create file fail: "+err.Error(), c)
			return
		}
		f.Write(file_data)
		file_data = nil

		//搜索boundary
		temp_data, reach_end, err := stream.ReadToBoundary(boundary, c.Request.Body, f)
		f.Close()
		if err != nil {
			engine.ERR("search boundary fail: "+err.Error(), c)
			return
		}
		if reach_end {
			break
		} else {
			copy(read_data[0:], temp_data)
			read_total = len(temp_data)
			continue
		}
	}
	//上传成功
	engine.OK("上传成功", "", c)
}
