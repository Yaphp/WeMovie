package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"wemovie/app/model"
	"wemovie/app/utils"
)

type ShareController struct {
	BaseController
}

// Add 分享文件
func (con ShareController) Add(c *gin.Context) {
	Fid := c.Query("id")
	Uid := c.Query("uid")
	Expire := c.Query("expire")
	Type := c.Query("type")

	// 生成分享文件的HashId
	hashId := utils.Md5(Fid + Uid + Expire + Type)

	// 判断是否已经分享过
	var share model.Share
	model.Db.Where("hash_id = ?", hashId).Find(&share)

	if share.Id > 0 {
		con.Success(c, "分享已存在")
		return
	}

	var password string

	// 如果type为2 则生成随机 4位 字母+数字 密码
	if Type == "2" {
		password = utils.RandomString(4)
	} else {
		password = ""
	}

	var shareFile model.Share

	shareFile.HashId = hashId
	shareFile.Fid = Fid
	shareFile.Uid = utils.StrToInt(Uid)
	shareFile.Expire = utils.StrToInt(Expire)
	shareFile.Type = utils.StrToInt(Type)
	shareFile.Password = password
	shareFile.CreatedAt = model.LocalTime(utils.GetDateTime())
	shareFile.UpdatedAt = model.LocalTime(utils.GetDateTime())

	err := model.Db.Create(&shareFile).Error

	if err != nil {
		con.Error(c, "分享失败")
		return
	}

	// 当前请求域名
	domain := c.Request.Host

	url := "https://" + domain + "#/share?file=" + hashId + "&password=" + password

	con.Success(c, url)
}

// Delete 删除分享文件
func (con ShareController) Delete(c *gin.Context) {

}

// Update 更新分享文件
func (con ShareController) Update(c *gin.Context) {

}

// Index 获取文件
func (con ShareController) Index(c *gin.Context) {
	fileHashId := c.Query("file")

	var share model.Share

	model.Db.Where("hash_id = ?", fileHashId).Find(&share)

	if share.Id <= 0 {
		con.Error(c, "分享不存在")
		return
	}

	// 判断分享是否过期
	expire := share.Expire
	createdAt := share.CreatedAt

	// 如果有效期为1则表示30天有效
	if expire == 1 {
		// 当前时间 > 30天 + 创建时间 把createdAt转换为时间戳
		layout := "2006-01-02 15:04:05"
		createdTime, _ := time.Parse(layout, string(createdAt))

		if time.Now().Unix() > createdTime.Unix()+30*24*60*60 {
			con.Error(c, "分享已过期")
			return
		}
	}

	// 判断密码是否正确
	if share.Password != "" {
		password := c.Query("password")
		if password != share.Password {
			con.Error(c, "密码错误")
			return
		}
	}

	// 查询结果
	var result = make(map[string]interface{})

	result["share"] = share

	// 分享用户信息
	var user model.User
	model.Db.Where("id = ?", share.Uid).Find(&user)
	result["user"] = user

	// 分享文件信息
	var file []model.File
	model.Db.Where("id in ?", strings.Split(share.Fid, ",")).Find(&file)
	result["files"] = file

	// 返回结果
	con.Success(c, result)
}

// Save 保存分享的文件
func (con ShareController) Save(c *gin.Context) {
	var param = make(map[string]interface{})

	json.NewDecoder(c.Request.Body).Decode(&param)

	Fid := param["id"]
	Uid := param["uid"]
	Pid := param["pid"]

	// 判断是否已经保存过
	var file model.File
	var total int64
	model.Db.Find(&file).Where("id in (?) and uid = ?", strings.Split(Fid.(string), ","), Uid).Count(&total)

	// 如果已经保存过则不再保存
	if total > 0 {
		con.Success(c, "文件已存在")
		return
	}

	// 获取文件信息
	var files []model.File
	model.Db.Where("id in (?)", strings.Split(Fid.(string), ",")).Find(&files)

	// 新文件记录
	var item []model.File

	// 文件基本信息
	var info = make(map[string]int)

	info["uid"] = int(Uid.(float64))
	info["pid"] = int(Pid.(float64))

	// 获取文件夹信息 及 子文件夹文件信息
	item = model.SaveDirAndFile(info, files, item)

	// 保存文件
	model.Db.CreateInBatches(item, len(item))

	con.Success(c, "保存成功")
}
