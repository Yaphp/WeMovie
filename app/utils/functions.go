package utils

/*
	工具函数集
*/

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// StrToInt 字符串转int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// StrToInt64 字符串转int64
func StrToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

// StrToFloat64 字符串转Float64
func StrToFloat64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
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

// RandomString 生成随机字符串
func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
