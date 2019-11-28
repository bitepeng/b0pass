package api

import (
	"b0pass/library/fileinfos"
	nustdbs "b0pass/library/nutsdbs"
	"b0pass/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"io"
	"log"
	"os"
)

// 执行文件上传处理
func Upload(r *ghttp.Request) {
	if err := r.ParseMultipartForm(32); err != nil {
		response.JSON(r, 201, err.Error())
	}
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer func() { _ = f.Close() }()
		name := gfile.Basename(h.Filename)
		size := h.Size
		// Get path
		pathSub :=r.GetPostString("path")
		nustdbs.DBs.SetData("files_path",pathSub)
		// Save path
		savePath := fileinfos.GetRootPath() + "/files/" +pathSub+"/"+ name
		log.Println(savePath)
		// Upload file
		file, err := gfile.Create(savePath)
		if err != nil {
			r.Response.Write(err)
			return
		}
		defer func() { _ = file.Close() }()
		if _, err := io.Copy(file, f); err != nil {
			response.JSON(r, 201, err.Error())
			return
		}
		response.JSON(r, 0, "ok", size)
	} else {
		response.JSON(r, 201, e.Error())
	}
}

// Uploadx 以小内存上传大文件
func Uploadx(r *ghttp.Request) {
	//Multipart Pipe
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		response.JSON(r, 201, err.Error())
	}
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer func() { _ = f.Close() }()
		name := gfile.Basename(h.Filename)

		//写入文件
		dst, err := os.OpenFile(
			fileinfos.GetRootPath()+"/files/"+name,
			os.O_WRONLY|os.O_CREATE, 0666,
		)
		defer func() { _ = dst.Close() }()
		if err != nil {
			response.JSON(r, 201, err.Error())
		}
		if _, err := io.Copy(dst, f); err != nil {
			response.JSON(r, 201, err.Error())
		}

		response.JSON(r, 0, "ok", name)
	} else {
		response.JSON(r, 201, e.Error())
	}
}

// Lists
func Lists(r *ghttp.Request) {
	fp := fileinfos.GetRootPath() + "/files/*"
	var ret []map[string]string
	ret = fileinfos.ListDirData(fp,"files")
	response.JSON(r, 0, "ok", ret)
}

// Delete
func Delete(r *ghttp.Request) {
	f := r.Get("f")
	fp := fileinfos.GetRootPath()
	filePath := fp + gconv.String(f)
	_ = os.RemoveAll(filePath)
	response.JSON(r, 0, "ok", filePath)
}

// Dump
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
