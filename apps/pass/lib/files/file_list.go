package files

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
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
	str := strings.Replace(dir, "\\", "/", -1)
	str = strings.TrimRight(str, "/")
	return str
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
	var imgs = []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".ico"}
	imgf := false
	for _, imge := range imgs {
		if strings.Index(strings.ToLower(f), imge) > 0 {
			imgf = true
		}
	}
	return imgf
}

// PathExists 文件是否存在
func PathExists(path string) (bool, fs.FileInfo, error) {
	fio, err := os.Stat(path)
	if err == nil {
		return true, fio, nil
	}
	if os.IsNotExist(err) {
		return false, nil, nil
	}
	return false, nil, err
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

// ValidUTF8 是否UTF8编码
func ValidUTF8(buf []byte) bool {
	nBytes := 0
	for i := 0; i < len(buf); i++ {
		if nBytes == 0 {
			if (buf[i] & 0x80) != 0 { //与操作之后不为0，说明首位为1
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1 //左移一位
					nBytes++     //记录字符共占几个字节
				}
				if nBytes < 2 || nBytes > 6 { //因为UTF8编码单字符最多不超过6个字节
					return false
				}
				nBytes-- //减掉首字节的一个计数
			}
		} else { //处理多字节字符
			if buf[i]&0xc0 != 0x80 { //判断多字节后面的字节是否是10开头
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}

// CheckFile 检查文件是否可读取
func CheckFile(filename string) (header string, canread bool) {
	file, _ := os.Open(filename)
	defer file.Close()
	buffer := make([]byte, 32)
	read, _ := file.Read(buffer)
	header = string(buffer[:read])
	canread = !strings.ContainsAny(header, "�")
	return
}

// GetCounts 文件数
func GetCounts(fp string) int {
	files, _ := filepath.Glob(fp + "/*")
	log.Println("GetCounts::", files)
	var i = 0
	for _, file := range files {
		fileInfo, _ := os.Stat(file)
		mfile := filepath.Base(file)
		if string(mfile[0]) == "." {
			continue
		}
		if fileInfo.IsDir() {
			i = i + GetCounts(file)
		} else {
			i++
		}
	}
	return i
}

// GetDirTree 文件列表
func GetDirTree(root, fp, fpSub, ftype string) []map[string]interface{} {
	files, _ := filepath.Glob(fp + "/*")
	log.Println("GetDirTree", fp, files)
	var ret []map[string]interface{}
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
		mext := strings.ToUpper(path.Ext(mfile))
		if fileInfo.IsDir() {
			mext = "DIR"
			mtype = "dir"
		}
		if mext == ".MP4" || mext == ".MOV" || mext == ".AVI" {
			mtype = "vod"
		}
		if mext == ".HTML" || mext == ".HTM" {
			mtype = "htm"
		}
		if mext == ".PDF" {
			mtype = "pdf"
		}
		if mext == "" {
			mext = "FILE"
			mtype = "file"
		}
		//index
		indexs++
		//path
		fpath := filepath.Join(fp, fpSub, mfile)
		//是否是文本类型文件
		header, canread := CheckFile(fpath)
		//path clear
		fpath = strings.Replace(fpath, root, "", 1)
		fpath = strings.ReplaceAll(fpath, "\\", "/")
		//map
		m := make(map[string]interface{})
		m["indexs"] = strconv.Itoa(indexs)
		m["name"] = mfile
		m["ext"] = mext
		m["size"] = strconv.Itoa(int(fileInfo.Size()))
		m["sizes"] = GetSize(uint64(fileInfo.Size()))
		m["date"] = fileInfo.ModTime().Format("01-02 15:04")
		m["path"] = fpath
		m["type"] = mtype
		m["canread"] = canread
		m["header"] = header
		if ftype != "" {
			if mtype == ftype {
				ret = append(ret, m)
			}
		} else {
			ret = append(ret, m)
		}
	}
	return ret
}

// GetData 文件内容
func GetData(fPath string) ([]byte, error) {
	fb, fio, err := PathExists(fPath)
	if err != nil {
		return []byte(err.Error()), err
	}
	if fb {
		if !fio.IsDir() {
			fInfo, _ := os.Stat(fPath)
			//是否是文本类型文件
			header, canread := CheckFile(fPath)
			if canread {
				//文件是否过大
				if int(fInfo.Size()) >= 2*1024*1024 {
					return []byte("+++++ [[Too Big]] Abstract: \n" + header + "\n+++++ ..."), nil
				}
				return ioutil.ReadFile(fPath)
			}
			return []byte("+++++ [[No Preview]] +++++"), nil
		} else {
			return nil, errors.New(fPath + " 是目录")
		}
	}
	return nil, errors.New(fPath + " 文件不存在")
}
