package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"wemovie/app/model"
	"wemovie/app/utils"
)

type FileController struct {
	BaseController
}

// Add 添加文件
func (con FileController) Add(c *gin.Context) {
	file := model.File{}

	err := c.ShouldBindJSON(&file)

	if err != nil {
		return
	}

	file.CreatedAt = model.LocalTime(utils.GetDateTime())

	fmt.Println(file)

	err = con.db().Create(&file).Error

	if err != nil {
		con.Error(c, "添加失败")
	} else {
		con.Success(c, "添加成功")
	}
}

// Delete 删除文件
func (con FileController) Delete(c *gin.Context) {
	id := c.Query("id")

	ids := strings.Split(id, ",")

	var file model.File

	err := con.db().Where("id in (?)", ids).Delete(&file).Error

	if err != nil {
		con.Error(c, "删除失败")
	} else {
		con.Success(c, "删除成功")
	}
}

// Update 更新文件
func (con FileController) Update(c *gin.Context) {
	file := model.File{}

	err := c.ShouldBindJSON(&file)
	if err != nil {
		return
	}

	fmt.Println(file)

	err = con.db().Where("id = ?", file.Id).Save(&file).Error

	if err != nil {
		con.Error(c, "更新失败")
	} else {
		con.Success(c, "更新成功")
	}
}

// Index 获取文件
func (con FileController) Index(c *gin.Context) {
	pageNo := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("limit", "10")
	keyword := c.Query("keyword")

	file := []model.File{}

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
		con.db().Where("filename like ?", "%"+keyword+"%").Find(&file).Count(&total)
		con.db().Where("filename like ?", "%"+keyword+"%").Offset((page - 1) * size).Order("id desc").Limit(size).Find(&file)
	} else {
		con.db().Find(&file).Count(&total)
		con.db().Offset((page - 1) * size).Order("id desc").Limit(size).Find(&file)
	}

	var result = make(map[string]interface{})

	result["total"] = total

	result["list"] = file

	con.Success(c, result)
}
