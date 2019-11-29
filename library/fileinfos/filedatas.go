package fileinfos

import (
	"fmt"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
)

// DataInit 从文件恢复为缓存
func Init(keys... string){
	fmt.Println(keys)
	for _,key:=range keys{
		data:=gfile.GetContents(cacheFile(key))
		gcache.Set(key,data,0)
	}
}

// DataSet 设置缓存
func Set(key,value string){
	gcache.Set(key,value,0)
	_ = gfile.PutContents(cacheFile(key),value)
}

// DataGet 读取缓存
func Get(key string) string{
	return gcache.Get(key).(string)
}

// cacheFile 缓存实例化文件
func cacheFile(key string) string{
	return GetRootPath()+"/tmp/data/"+key+".txt"
}