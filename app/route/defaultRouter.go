package route

import (
	"github.com/gin-gonic/gin"
)

func DefaultRouter(r *gin.Engine) {
	router := r.Group("/")

	//默认跳转至dist文件夹下的index.html
	router.GET("", func(c *gin.Context) {
		c.HTML(200, "./dist/index.html", gin.H{})
	})
}
