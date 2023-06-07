package main

import (
	"b0pass/boot"
	_ "b0pass/boot"
	"b0pass/library/openurl"
	_ "b0pass/router"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"
)

func main() {

	//处理命令行参数
	boot.ExecArgs()
	fmt.Printf("[ServerUrl] http://127.0.0.1:%d\n",boot.ServPort)
	fmt.Printf("[Work-Path] %s\n",boot.PathRoot)

	//是否开启GUI模式
	//判断是否安装谷歌浏览器
	ChromeExe := lorca.ChromeExecutable()
	if ChromeExe != "" {
		//打开UI界面
		execUI()
	} else {
		//打开浏览器
		go func() {
			time.Sleep(1000 * time.Millisecond)
			_ = openurl.Open("http://127.0.0.1:" + strconv.Itoa(boot.ServPort))
		}()
		g.Wait()
	}
}

 func execUI() {
	// Wait Server Run
	time.Sleep(3 * time.Second)

	// Cli Args
	var args []string
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	if runtime.GOOS == "windows" {
		args = append(args, "-ldflags '-H windowsgui'", "--remote-allow-origins=*")
	}

	// New Lorca UI
	ui, err := lorca.New(
		`data:text/html,
		<html><head><title>B0App</title></head></html>`,
		"", 360, 640, args...,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = ui.Close()
	}()

	// Load url
	_ = ui.Load(fmt.Sprintf(
		"http://%s",
		"127.0.0.1:"+g.Config().GetString("setting.port")),
	)

	// Wait until the interrupt signal arrives
	// or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	// Close UI
	log.Println("exiting...")
	_ = g.Server().Shutdown()
}