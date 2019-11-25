package nustdbs

import (
	"fmt"
	"testing"
)


func TestDBClient(t *testing.T) {

	//最后关闭连接
	defer DBs.db.Close()

	//写入数据
	keys:="kkkk"
	value:="vvvv"
	DBs.SetData(keys, value)
	keys="key1"
	value="vvvv11"
	DBs.SetData(keys, value)
	keys="key2"
	value="vvvv22"
	DBs.SetData(keys, value)

	//删除数据
	DBs.DelData("key1")

	//查询数据
	data:=DBs.GetData(keys)
	fmt.Println(keys)
	fmt.Println(data)



	//查询全部
	datas:=DBs.GetDatas("",100)
	fmt.Println(datas)

	//查询predfix全部
	datas=DBs.GetDatas("key",100)
	fmt.Println(datas)

}
