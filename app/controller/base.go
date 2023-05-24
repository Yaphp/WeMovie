package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"wemovie/app/model"
)

type BaseController struct{}

// Success 成功返回
func (con BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

// SuccessWithCount 返回带总数的数据
func (con BaseController) SuccessWithCount(c *gin.Context, data interface{}, count int64) {
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"data":  data,
		"count": count,
	})
}

// Error 错误返回
func (con BaseController) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
	})
}

// CheckLogin 验证登录
func (con BaseController) CheckLogin(c *gin.Context) bool {
	return true
}

// db 获取数据库连接
func (con BaseController) db() *gorm.DB {
	return model.Db
}

// Params 获取json参数 (int类型会丢失)
func (con BaseController) Params(c *gin.Context) map[string]string {
	var params map[string]string
	c.BindJSON(&params)
	//fmt.Println(params)
	return params
}
