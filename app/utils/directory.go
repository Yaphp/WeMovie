package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckFileIsExist(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func CreateDir(dir string) bool {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("创建文件夹失败", err)
		return false
	}
	return true
}

// GetRootPath 获取根目录
func GetRootPath() string {
	exePath, _ := os.Executable()
	rootPath := filepath.Dir(exePath)
	//println(currentPath)
	return rootPath
}
