package files

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// 目录层级
var level int

func init() {
	level = 0
}

// FileNode 节点结构
type FileNode struct {
	Name      string      `json:"title"`
	Path      string      `json:"path"`
	Spread    bool        `json:"spread"`
	FileNodes []*FileNode `json:"children"`
}

// NodeTree 目录树数据
func NodeTree(rootpath, f string) FileNode {
	root := FileNode{rootpath, "/", true, []*FileNode{}}
	fileInfo, _ := os.Lstat(rootpath)
	walk(rootpath, rootpath, f, fileInfo, &root)
	//data, _ := json.Marshal(root)
	//fmt.Printf("%s", data)
	return root
}

// 添加节点
func NodeAdd(fPath string) error {
	dir, file := filepath.Split(fPath)
	log.Println("::NodeAdd-> dir=", dir, "file=", file)
	if err := os.MkdirAll(dir, 0666); err != nil {
		return err
	}
	if file != "" {
		return os.WriteFile(fPath, []byte(""), 0666)
	}
	return nil
}

// 重命名节点
func NodeRename(fPath, nPath string) error {
	return os.Rename(fPath, nPath)
}

// 删除节点
func NodeRemove(fPath string) error {
	return os.Remove(fPath)
}

// walk 递归查询所有目录节点
func walk(root, path, f string, info os.FileInfo, node *FileNode) {
	// 列出当前目录下的所有目录、文件
	files := listFiles(path)
	// 遍历这些文件
	for _, filename := range files {
		// 拼接全路径
		fpath := filepath.Join(path, filename)
		// 构造文件结构
		fio, _ := os.Lstat(fpath)
		//node.FileNodes
		// 如果遍历的当前文件是个目录，则进入该目录进行递归
		if fio.IsDir() {
			// 清理路径为相对路径
			fpath_ := strings.Replace(fpath, root, "", 1)
			fpath_ = strings.ReplaceAll(fpath_, "\\", "/")
			//是否默认展开
			level++
			Spread := false
			if level < 8 || strings.Contains(f, fpath_) || fpath_ == f {
				Spread = true
			}
			// 将当前文件作为子节点添加到目录下
			child := FileNode{filename, fpath_, Spread, []*FileNode{}}
			node.FileNodes = append(node.FileNodes, &child)
			walk(root, fpath, f, fio, &child)
		}
	}
}

// listFiles 列出所有文件名
func listFiles(dirname string) []string {
	f, _ := os.Open(dirname)
	names, _ := f.Readdirnames(-1)
	f.Close()
	sort.Strings(names)
	return names
}
