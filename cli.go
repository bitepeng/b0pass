package main

import (
	_ "b0pass/boot"
	"b0pass/library/ipaddress"
	_ "b0pass/router"
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func main() {
	/*
		Wait Server
	*/
	fmt.Println(ipaddress.GetIP())
	fmt.Println(g.Config().GetInt("setting.port"))
	g.Wait()
}
