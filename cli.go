package main

import (
	"b0pass/boot"
	_ "b0pass/boot"
	"b0pass/library/ipaddress"
	"b0pass/library/openurl"
	_ "b0pass/router"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"strconv"
	"time"
)

func main() {
	//Cli Args
	boot.ExecArgs()
	ipArr,_:=ipaddress.GetIP()
	fmt.Printf("[ServerUrl] http://127.0.0.1:%d\n",boot.ServPort)
	fmt.Printf("[IPlistArr] %v\n",ipArr)
	fmt.Printf("[Work-Path] %s\n",boot.PathRoot)
	//Open Urls
	go func() {
		time.Sleep(4000 * time.Millisecond)
		_ = openurl.Open("http://127.0.0.1:" + strconv.Itoa(boot.ServPort))
	}()
	g.Wait()
}
