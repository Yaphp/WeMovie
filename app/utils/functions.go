package utils

/*
	工具函数集
*/

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

// StrToInt 字符串转int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// IntToStr 数字转字符串
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// Md5 md5加密
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// GetDate 获取当前日期
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

// GetDateTime 获取当前日期时间
func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
