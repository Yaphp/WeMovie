package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"web_standard/app/model"
)

type BaseController struct{}

// 成功返回
func (con BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// 错误返回
func (con BaseController) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  msg,
	})
}

// CheckLogin 验证登录
func (con BaseController) CheckLogin(c *gin.Context) bool {
	return true
}

// 获取db链接
func (con BaseController) db() *gorm.DB {
	return model.Db
}

// 获取json参数 (int类型会丢失)
func (con BaseController) Params(c *gin.Context) map[string]string {
	var params map[string]string
	c.BindJSON(&params)
	//fmt.Println(params)
	return params
}
