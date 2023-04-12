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

	c.ShouldBindJSON(&user)

	user.Password = utils.Md5(user.Password)

	user.CreatedAt = model.LocalTime(utils.GetDateTime())

	fmt.Println(user)

	err := con.db().Create(&user).Error

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

	err := con.db().Where("id in (?)", ids).Delete(&user).Error

	if err != nil {
		con.Error(c, "删除失败")
	} else {
		con.Success(c, "删除成功")
	}
}

// Update 更新用户
func (con UserController) Update(c *gin.Context) {
	user := model.User{}

	c.ShouldBindJSON(&user)

	fmt.Println(user)

	var err error

	if user.Password == "" {
		err = con.db().Omit("password").Where("id = ?", user.Id).Save(&user).Error
	} else {
		user.Password = utils.Md5(user.Password)
		err = con.db().Where("id = ?", user.Id).Save(&user).Error
	}

	if err != nil {
		con.Error(c, "更新失败")
	} else {
		con.Success(c, "更新成功")
	}
}

// Index 获取用户
func (con UserController) Index(c *gin.Context) {
	pageNo := c.DefaultQuery("pageNo", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	keyword := c.Query("keyword")

	user := []model.User{}

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
		con.db().Where("username like ?", "%"+keyword+"%").Find(&user).Count(&total)
		con.db().Where("username like ?", "%"+keyword+"%").Offset((page - 1) * size).Order("id desc").Limit(size).Find(&user)
	} else {
		con.db().Find(&user).Count(&total)
		con.db().Offset((page - 1) * size).Order("id desc").Limit(size).Find(&user)
	}

	var result = make(map[string]interface{})

	result["total"] = total

	result["list"] = user

	con.Success(c, result)
}
