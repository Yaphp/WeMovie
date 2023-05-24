package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"wemovie/app/model"
	"wemovie/app/utils"
)

type UserController struct {
	BaseController
}

// Add 添加用户
func (con UserController) Add(c *gin.Context) {
	user := model.User{}

	err := c.ShouldBindJSON(&user)

	if err != nil {
		con.Error(c, "参数解析失败")
		return
	}

	// md5加密password + salt
	user.Password = utils.Md5(user.Password + "810169")

	user.CreatedAt = model.LocalTime(utils.GetDateTime())

	fmt.Println(user)

	err = model.Db.Create(&user).Error

	if err != nil {
		con.Error(c, "添加失败")
	} else {
		con.Success(c, "添加成功")
	}
}

// Delete 删除用户
func (con UserController) Delete(c *gin.Context) {
	id := c.Query("id")

	ids := strings.Split(id, ",")

	var user model.User

	err := model.Db.Where("id in (?)", ids).Delete(&user).Error

	if err != nil {
		con.Error(c, "删除失败")
	} else {
		con.Success(c, "删除成功")
	}
}

// Update 更新用户
func (con UserController) Update(c *gin.Context) {
	user := model.User{}

	err := c.ShouldBindJSON(&user)

	if err != nil {
		con.Error(c, "参数解析失败")
		return
	}

	fmt.Println(user)

	if user.Password == "" {
		err = model.Db.Omit("password").Where("id = ?", user.Id).Save(&user).Error
	} else {
		user.Password = utils.Md5(user.Password + "810169")
		err = model.Db.Where("id = ?", user.Id).Save(&user).Error
	}

	if err != nil {
		con.Error(c, "更新失败")
	} else {
		// 统计使用的内存总量
		var result []int64
		var total int64
		model.Db.Model(&model.File{}).Where("uid = ?", user.Id).Pluck("Size", &result)
		for _, v := range result {
			total += v
		}
		con.SuccessWithCount(c, user, total)
	}
}

// Index 获取用户
func (con UserController) Index(c *gin.Context) {
	pageNo := c.DefaultQuery("pageNo", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	keyword := c.Query("keyword")

	var user []model.User

	var total int64

	// 把params["page"]转换成int
	page := utils.StrToInt(pageNo)
	size := utils.StrToInt(pageSize)

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	if keyword != "" {
		model.Db.Where("username like ?", "%"+keyword+"%").Find(&user).Count(&total)
		model.Db.Where("username like ?", "%"+keyword+"%").Offset((page - 1) * size).Order("id desc").Limit(size).Find(&user)
	} else {
		model.Db.Find(&user).Count(&total)
		model.Db.Offset((page - 1) * size).Order("id desc").Limit(size).Find(&user)
	}

	con.SuccessWithCount(c, user, total)
}
