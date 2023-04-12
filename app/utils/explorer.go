package utils

/*
if runtime.GOOS == "windows" {
	// 打开浏览器
	cmd := exec.Command("cmd", "/c start http://127.0.0.1")

	// 启动服务
	res := cmd.Start()

	fmt.Println(res)
	fmt.Println("启动浏览器成功")
}
*/

import (
	"fmt"
	"os/exec"
	"runtime"
)

var commands = map[string]string{
	"windows": "cmd",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Openurl(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	if runtime.GOOS == "windows" {
		return exec.Command(run, "/c", "start", uri).Start()
	} else {
		return exec.Command(run, uri).Start()
	}
}
