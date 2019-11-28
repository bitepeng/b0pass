package main


import (
	nustdbs "b0pass/library/nutsdbs"
	"fmt"
)

func main() {
	testNutsdb()
}

func testNutsdb(){
	nustdbs.DBs.SetData("key","ccccccccc")
	data:=nustdbs.DBs.GetDatas("",100)
	fmt.Println(data)
}
