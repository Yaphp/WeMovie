package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"web_standard/app/utils"
)

type UploadController struct {
	BaseController
}

// Add 上传图片
func (con UploadController) Add(c *gin.Context) {
	file, _ := c.FormFile("file")

	// 获取当前日期
	date := time.Now().Format("20060102")

	//使用md5生成文件名
	fileName := fmt.Sprintf("%x", md5.Sum([]byte(file.Filename+time.Now().String())))

	// 获取文件后缀
	fileSuffix := file.Filename[strings.LastIndex(file.Filename, "."):]

	// 检测文件夹是否存在 不存在则创建
	if !utils.CheckFileIsExist("./uploads/" + date) {
		utils.CreateDir("./uploads/" + date)
	}

	//上传文件至指定目录
	err := c.SaveUploadedFile(file, "./uploads/"+date+"/"+fileName+fileSuffix)

	if err != nil {
		println(err)
		con.Error(c, "上传失败")
		return
	}

	// 返回结果
	con.Success(c, "/uploads/"+date+"/"+fileName+fileSuffix)
}
