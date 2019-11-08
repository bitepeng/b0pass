package nustdb

import (
	"fmt"
	"log"
	"testing"
)

func TestDBClient(t *testing.T) {

	//最后关闭连接
	defer DBs.db.Close()

	//写入数据
	keys:="kkkk"
	value:="vvvv"
	if err := DBs.SetData(keys, value);err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println(keys)
	}

	//查询数据
	if data,err:=DBs.GetData(keys);err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println(data)
	}
}