package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
