package main

import (
	"b0pass/boot"
	_ "b0pass/boot"
	_ "b0pass/router"
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func main() {
	/*
		Wait Server
	*/
	boot.ExecArgs()
	fmt.Printf("[ServerUrl] http://127.0.0.1:%d\n",boot.ServPort)
	fmt.Printf("[Work-Path] %s\n",boot.PathRoot)
	g.Wait()
}
