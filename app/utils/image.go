package utils

import (
	"fmt"
	"github.com/disintegration/imaging"
	"strings"
)

// GetThumbnailImage 生成缩略图
func GetThumbnailImage(file string) string {
	src, err := imaging.Open(file)
	if err != nil {
		fmt.Printf("failed to open image: %v\n", err)
	}

	// 生成缩略图
	dst := imaging.Thumbnail(src, 100, 100, imaging.Lanczos)

	// 获取图片名称
	fileName := file[strings.LastIndex(file, "/")+1:]

	// 替换名称
	replaceName := strings.Replace(file, fileName, "thumbnail_"+fileName, 1)

	// 保存
	err = imaging.Save(dst, replaceName)
	if err != nil {
		fmt.Printf("failed to save image: %v\n", err)
		return ""
	}

	// 返回缩略图地址
	return "thumbnail_" + fileName
}
