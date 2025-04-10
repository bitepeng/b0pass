package app

import (
	"b0go/apps/pass/lib/files"
	"b0go/apps/pass/lib/stream"
	"b0go/core/engine"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//----------------------RO-----------------------//

/**
* NodeTree 目录树结构
* @param f 目录路径
 */
func NodeTree(c *gin.Context) {
	f := c.Query("f")
	listMap := files.NodeTree(config.Path, f)
	engine.OK("OK", listMap, c)
}

/**
* FileCount 文件统计
 */
func FileCount(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	listMap := files.GetCounts(RootPath)
	engine.OK("OK", listMap, c)
}

/**
* FileList 文件列表
* @param t 文件类型，可选值：file, dir, all
 */
func FileList(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	f := c.Query("f")
	t := c.DefaultQuery("t", "")
	if f == "/" || !filepath.IsAbs(f) {
		f = cleanPathJoin(RootPath, f)
	}
	listMap := files.GetDirTree(RootPath, f, "", t)
	engine.OK("OK", listMap, c)
}

// FileContent 文件内容
func FileContent(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	f := c.Query("f")
	if f == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	f = cleanPathJoin(RootPath, f)
	data, err := files.GetData(f)
	if err != nil {
		engine.ERR(err.Error(), c)
		return
	}
	engine.OK("OK", string(data), c)
}

// FileDownload 文件下载
func FileDownload(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	f := c.Query("f")
	if f == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	f = cleanPathJoin(RootPath, f)
	//获取文件的名称
	fileName := path.Base(f)

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	//fmt.Println(filePath)
	c.File(f)
}

//----------------------RW-----------------------//

/**
* NodeTree 重命名目录
* @param f 目录路径
* @param n 新名称
 */
func NodeRename(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	// 获取路径和新名称
	f := cleanPath(c.Query("f"))
	n := cleanPath(c.Query("n"))
	// 检查路径和新名称是否为空
	if f == "" || n == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	f = cleanPathJoin(RootPath, f)
	n = cleanPathJoin(RootPath, n)
	log.Printf("::NodeRename:: %s %s", f, n)

	// 检查锁定上传目录
	if !checkLockUploadDir(n, c) {
		return
	}

	// 调用重命名函数
	err := files.NodeRename(f, n)
	if err != nil {
		engine.ERR(err.Error(), c)
		return
	}

	// 重命名成功
	engine.OK("OK", nil, c)
}

/**
* NodeRemove 删除目录
* @param f 目录路径
 */
func NodeRemove(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	f := cleanPath(c.Query("f"))
	if f == "" {
		engine.ERR("路径不能为空", c)
		return
	}

	f = cleanPathJoin(RootPath, f)
	log.Printf("::NodeRemove:: %s", f)

	// 检查锁定上传目录
	if !checkLockUploadDir(f, c) {
		return
	}

	err := files.NodeRemove(f)
	if err != nil {
		engine.ERR(err.Error(), c)
		return
	}
	engine.OK("OK", nil, c)
}

// NodeAdd 添加目录
func NodeAdd(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	f := cleanPath(c.Query("f"))
	if f == "" {
		engine.ERR("路径不能为空", c)
		return
	}
	f = cleanPathJoin(RootPath, f)
	log.Printf("::NodeAdd:: %s", f)

	// 检查锁定上传目录
	if !checkLockUploadDir(f, c) {
		return
	}

	err := files.NodeAdd(f + "/")
	if err != nil {
		engine.ERR(err.Error(), c)
	}
	engine.OK("OK", nil, c)
}

// FileUpload 上传文件
func FileUpload(c *gin.Context) {
	RootPath := cleanPath(config.Path)
	lens, _ := strconv.Atoi(c.Request.Header["Content-Length"][0])
	log.Println("FileUpload::::", lens)
	f := c.DefaultQuery("f", "/")
	f = cleanPathJoin(RootPath, f)
	//检查是否锁定上传目录
	if config.LockUploadDir != "" {
		if !strings.Contains(f, "/"+config.LockUploadDir) {
			f = config.Path + "/" + config.LockUploadDir + "/"
			engine.ERR("上传目录被锁定为："+config.LockUploadDir, c)
		}
	}
	files.NodeAdd(f)
	if lens > 4096 {
		FileUploadBig(f, c)
	} else {
		FileUploadTiny(f, c)
	}
}

// FileUploadTiny 小文件上传
func FileUploadTiny(RootPath string, c *gin.Context) {
	log.Println("FileUploadTiny:::", RootPath)
	//上传文件
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, RootPath+file.Filename)
	//上传成功
	engine.OK("上传成功", "", c)
}

// FileUploadBig 大文件上传
func FileUploadBig(RootPath string, c *gin.Context) {
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
		fn := cleanPathJoin(RootPath, file_header.FileName)
		log.Printf("save file: [%s]\n", fn)
		f, err := os.Create(fn)
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
