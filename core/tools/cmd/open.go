package cmd

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// commands 执行程序
var commands = map[string]string{
	"windows": `cmd /c start `,
	"darwin":  "open ",
	"linux":   "xdg-open ", //eog -w
}

// divisions 路径分隔符
var divisions = map[string]string{
	"windows": "\\",
	"darwin":  "/",
	"linux":   "/",
}

// Open 打开浏览器
func Open(uri string) error {
	//runtime.GOOS
	run, ok := commands[runtime.GOOS]
	if !ok {
		log.Printf("don't know how to open things on %s platform", runtime.GOOS)
	}
	//uri divisions
	div := divisions[runtime.GOOS]
	uri = strings.ReplaceAll(uri, "\\", div)
	uri = strings.ReplaceAll(uri, "/", div)
	//exec.Command
	run = run + uri
	cmds := strings.Split(run, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	//cmd.Start
	log.Println("[CommandAs]", cmds)
	//windows cmd不出现黑框
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Start()
}
