package main

import (
	_ "b0pass/boot"
	"b0pass/library/ipaddress"
	_ "b0pass/router"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gres"
)

func main() {
	/*
		Wait Server
	*/
	gres.Dump()
	fmt.Println(ipaddress.GetIP())
	g.Wait()
}
