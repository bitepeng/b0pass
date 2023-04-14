package files

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// GetBinPath 获取当前可执行文件的路径
func GetBinPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//str := strings.Replace(dir, "\\", "/", -1)
	//str := strings.TrimRight(str, "/")
	return dir
}

// GetCodePath 获取当前代码文件路径
func GetCodePath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Can not get current file info")
	}
	return filepath.Dir(file)
}

// IfImage 根据文件名判断是否是图片
func IfImage(f string) bool {
	imgs := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".ico"}
	imgf := false
	for _, imge := range imgs {
		if strings.Index(strings.ToLower(f), imge) > 0 {
			imgf = true
		}
	}
	return imgf
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetSize 显示友好的文件大小
func GetSize(fileBytes uint64) string {
	var (
		units []string
		size  string
		i     int
	)
	units = []string{"B", "K", "M", "G", "T", "P"}
	i = 0
	for {
		i++
		fileBytes = fileBytes / 1024
		if fileBytes < 1024 {
			size = fmt.Sprintf("%d", fileBytes) + units[i]
			break
		}
	}
	return size
}

// ListDirData 文件列表
// TODO: 安全检查，不能超过根目录限制
func ListDirData(fp, fpSub string) []map[string]string {
	log.Println("ListDirData", fp)
	files, _ := filepath.Glob(fp)
	var ret []map[string]string
	var indexs = 0
	for _, file := range files {
		fileInfo, _ := os.Stat(file)
		//filename
		mfile := filepath.Base(file)
		if string(mfile[0]) == "." {
			continue
		}
		//filetype
		mtype := "file"
		if IfImage(mfile) {
			mtype = "img"
		}
		//fileext strings.ToUpper()
		mext := path.Ext(mfile)
		if fileInfo.IsDir() {
			mext = "dir"
			mtype = "dir"
		}
		if mext == "" {
			mext = "file"
			mtype = "file"
		}
		//index
		indexs++
		//map
		m := make(map[string]string)
		m["name"] = mfile
		m["ext"] = mext
		m["size"] = strconv.Itoa(int(fileInfo.Size()))
		m["sizes"] = GetSize(uint64(fileInfo.Size()))
		m["date"] = fileInfo.ModTime().Format("01-02")
		m["path"] = fpSub + "/" + mfile
		m["type"] = mtype
		m["indexs"] = strconv.Itoa(indexs)
		ret = append(ret, m)
	}
	return ret
}
