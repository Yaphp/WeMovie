package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"os"
	"strings"
	"time"
	"wemovie/app/model"
	"wemovie/app/utils"
)

type UploadController struct {
	BaseController
}

// Index 上传图片
func (con UploadController) Index(c *gin.Context) {
	// 获取上传的文件
	f, _ := c.FormFile("file")

	// 获取上传的文件名
	uploadFileName := c.PostForm("name")

	// 如果没有传name参数则使用文件名
	if uploadFileName == "" {
		uploadFileName = f.Filename
	}

	// 获取当前日期
	date := time.Now().Format("20060102")

	// 使用uuid生成文件名
	fileName := uuid.New().String()

	// 获取文件后缀
	fileSuffix := uploadFileName[strings.LastIndex(uploadFileName, "."):]

	// 检测文件夹是否存在 不存在则创建
	if !utils.CheckFileIsExist(utils.GetRootPath() + "/dist/uploads/" + date) {
		utils.CreateDir(utils.GetRootPath() + "/dist/uploads/" + date)
	}

	//上传文件至指定目录
	err := c.SaveUploadedFile(f, utils.GetRootPath()+"/dist/uploads/"+date+"/"+fileName+fileSuffix)

	if err != nil {
		println(err)
		con.Error(c, "上传失败")
		return
	}

	// 根据save参数判断是否需要存入数据库
	save := c.PostForm("save")

	if save == "false" {
		con.Success(c, "/uploads/"+date+"/"+fileName+fileSuffix)
		return
	}

	// 保存到数据库
	var file model.File

	// 判断是否为图片
	if strings.Contains(c.PostForm("type"), "image") {
		// 生成缩略图
		thumb := utils.GetThumbnailImage(utils.GetRootPath() + "/dist/uploads/" + date + "/" + fileName + fileSuffix)
		file.Thumb = "/uploads/" + date + "/" + thumb
	} else {
		file.Thumb = ""
	}

	// 保存到数据库
	file.Pid = utils.StrToInt(c.PostForm("pid"))
	file.Uid = utils.StrToInt(c.PostForm("uid"))
	file.Name = c.PostForm("name")
	file.Type = c.PostForm("type")
	file.Size = utils.StrToFloat64(c.PostForm("size"))
	file.Path = "/uploads/" + date + "/" + fileName + fileSuffix
	file.CreatedAt = model.LocalTime(utils.GetDateTime())
	file.UpdatedAt = model.LocalTime(utils.GetDateTime())

	// 保存到数据库
	model.Db.Create(&file)

	// 判断保存是否成功
	if file.Id > 0 {
		con.Success(c, "/uploads/"+date+"/"+fileName+fileSuffix)
	} else {
		con.Error(c, "上传失败")
	}
}

// Chunk 分片上传
func (con UploadController) Chunk(c *gin.Context) {
	name := c.PostForm("name")
	path := c.PostForm("path")

	if c.PostForm("merge") == "true" {
		var file model.File
		if strings.Contains(c.PostForm("type"), "image") {
			// 生成缩略图
			thumbName := utils.GetThumbnailImage(path)
			uuidName := strings.Split(thumbName, "_")[1]
			replace := strings.Replace(path, uuidName, thumbName, 1)
			file.Thumb = replace
		}

		file.Pid = utils.StrToInt(c.PostForm("pid"))
		file.Uid = utils.StrToInt(c.PostForm("uid"))
		file.Name = name
		file.Type = c.PostForm("type")
		file.Size = utils.StrToFloat64(c.PostForm("size"))
		file.Path = strings.Split(path, "/dist")[1]
		file.CreatedAt = model.LocalTime(utils.GetDateTime())
		file.UpdatedAt = model.LocalTime(utils.GetDateTime())

		// 保存到数据库
		model.Db.Create(&file)

		// 判断保存是否成功
		if file.Id > 0 {
			con.Success(c, "上传成功")
		} else {
			con.Error(c, "合并文件失败")
		}
		return
	}

	f, _ := c.FormFile("chunk")
	chunkNumber := c.PostForm("chunk_number")

	// 无路径时生成路径
	if path == "" {
		date := time.Now().Format("20060102")

		// 生成文件uuid
		uuidName := uuid.New().String()

		// 保存路径
		path = utils.GetRootPath() + "/dist/uploads/" + date

		// 判断文件夹是否存在 不存在则创建
		if !utils.CheckFileIsExist(path) {
			createResutl := utils.CreateDir(path)

			if !createResutl {
				con.Error(c, "文件夹创建失败")
				return
			}
		}

		// 保存文件路径
		path = path + "/" + uuidName + name[strings.LastIndex(name, "."):]
	}

	// 计算偏移量
	offset := utils.StrToInt64(chunkNumber) * 1024 * 1024 * 10

	// 打开文件
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	// 移动指针到偏移量位置
	if _, err = file.Seek(offset, 0); err != nil {
		panic(err)
	}

	// 把f转为Reader file对象
	fReader, err := f.Open()
	if err != nil {
		panic(err)
	}

	// 把f的内容拷贝到file
	if _, err = io.Copy(file, fReader); err != nil {
		panic(err)
	}

	// 关闭文件
	file.Close()

	// 关闭文件
	fReader.Close()

	con.Success(c, path)
}
