package utils

import (
	"fmt"
	"os"
)

func CheckFileIsExist(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func CreateDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("创建文件夹失败", err)
	}
}
