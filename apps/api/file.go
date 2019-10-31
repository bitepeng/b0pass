package api

import (
	"b0pass/library/fileinfos"
	"b0pass/library/response"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"io"
	"os"
)

// 执行文件上传处理
func Upload(r *ghttp.Request) {
	if err:=r.ParseMultipartForm(32);err!=nil{
		response.JSON(r, 201, err.Error())
	}
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer func() {
			if err := f.Close();err!=nil{
				fmt.Println("f.Close() err: ",err)
			}
		}()
		name := gfile.Basename(h.Filename)
		buffer := make([]byte, h.Size)
		_, _ = f.Read(buffer)
		if err := gfile.PutBytes(fileinfos.GetRootPath()+"/files/"+name, buffer); err != nil {
			response.JSON(r, 201, err.Error(), name)
		}
		buffer=nil
		response.JSON(r, 0, "ok", buffer)
	} else {
		response.JSON(r, 201, e.Error())
	}
}

// Uploadx 以小内存上传大文件
func Uploadx(r *ghttp.Request){
	//Multipart Pipe
	if err:=r.ParseMultipartForm(32<<20);err!=nil{
		response.JSON(r, 201, err.Error())
	}
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer func() {_ = f.Close()}()
		name := gfile.Basename(h.Filename)

		//写入文件
		dst, err := os.OpenFile(
			fileinfos.GetRootPath()+"/files/"+name,
			os.O_WRONLY|os.O_CREATE, 0666,
			)
		defer func() {dst.Close()}()
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
	ret = fileinfos.ListDirData(fp)
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
