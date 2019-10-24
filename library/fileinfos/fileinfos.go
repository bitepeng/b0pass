package fileinfos

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取当前可执行文件的路径
func GetBinPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Replace(dir, "\\", "/", -1)
	str = strings.TrimRight(str, "/")
	return str
}

// 获取当前代码文件路径
func GetCodePath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Can not get current file info")
	}
	return filepath.Dir(file)
}

// getRootPath
func GetRootPath() string {
	var fp string
	//fmt.Println("os.Args >>>>> ",os.Args[0][0],os.Args)
	/*if os.Args[0][0]==47 {//exe 47==/
		fp=strings.Replace(GetCodePath(),"/library/fileinfos","",-1)
	}else{
		fp=GetBinPath()
	}*/
	fp = GetBinPath()
	return fp
}

// 根据文件名判断是否是图片
func IfImage(f string) bool {
	var imgs = []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".ico"}
	imgf := false
	for _, imge := range imgs {
		if strings.Index(strings.ToLower(f), imge) > 0 {
			imgf = true
		}
	}
	return imgf
}

// 判断文件夹是否存在
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
