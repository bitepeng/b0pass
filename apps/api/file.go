package api

import (
	"b0pass/library/fileinfos"
	"b0pass/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// 执行文件上传处理，上传到系统临时目录 /tmp
func Upload(r *ghttp.Request) {
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer func() { _ = f.Close() }()
		name := gfile.Basename(h.Filename)
		buffer := make([]byte, h.Size)
		_, _ = f.Read(buffer)
		if err := gfile.PutBytes(fileinfos.GetRootPath()+"/files/"+name, buffer); err != nil {
			response.JSON(r, 201, err.Error(), name)
		}
		response.JSON(r, 0, "ok", name)
	} else {
		response.JSON(r, 201, e.Error())
	}
}

// Lists
func Lists(r *ghttp.Request) {
	fp := fileinfos.GetRootPath() + "/files/*"
	files, _ := filepath.Glob(fp)
	var ret []map[string]string
	for _, file := range files {
		fileInfo, _ := os.Stat(file)
		//filename
		mfile := filepath.Base(file)
		if string(mfile[0]) == "." {
			continue
		}
		//filetype
		mtype := "file"
		if fileinfos.IfImage(mfile) {
			mtype = "img"
		}
		//fileext
		mext := strings.ToUpper(path.Ext(mfile))
		if fileInfo.IsDir() {
			mext = "目录"
		}
		//map
		m := make(map[string]string)
		m["name"] = mfile
		m["ext"] = mext
		m["size"] = strconv.Itoa(int(fileInfo.Size()))
		m["date"] = fileInfo.ModTime().Format("01-02")
		m["path"] = "/files/" + mfile
		m["type"] = mtype
		ret = append(ret, m)
	}
	response.JSON(r, 0, "ok", ret)
}

// Delete
func Delete(r *ghttp.Request) {
	f := r.Get("f")
	fp := fileinfos.GetRootPath()
	filePath := fp + gconv.String(f)
	_ = os.Remove(filePath)
	response.JSON(r, 0, "ok", filePath)
}

func Dump(r *ghttp.Request) {
	filePath := os.Args[0]
	response.JSON(r, 0, "ok", filePath)
}

// 展示文件上传页面
func UploadShow(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>上传文件</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/api/upload" method="post">
                <input type="file" name="upload-file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}
