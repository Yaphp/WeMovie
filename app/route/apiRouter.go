package route

import (
	"github.com/gin-gonic/gin"
	"wemovie/app/controller"
	"wemovie/app/model"
)

// Auth 验证token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		//println(token)
		if token == "" {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "未登录或非法请求",
			})
			c.Abort()
			return
		}

		user := model.User{}
		rows := model.Db.Where("token = ?", token).Find(&user).RowsAffected
		//println(rows)
		if rows == 0 {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "未登录或登录已过期",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// ApiRouter api路由
func ApiRouter(r *gin.Engine) {
	routerWithOutLogin := r.Group("/api")
	{
		routerWithOutLogin.POST("/register", controller.LoginController{}.Add)
		routerWithOutLogin.POST("/login", controller.LoginController{}.Index)
	}

	routerWithLogin := r.Group("/api").Use(Auth())
	{
		routerWithLogin.POST("/upload", controller.UploadController{}.Add)

		routerWithLogin.POST("/user", controller.UserController{}.Add)
		routerWithLogin.DELETE("/user", controller.UserController{}.Delete)
		routerWithLogin.PUT("/user", controller.UserController{}.Update)
		routerWithLogin.GET("/user", controller.UserController{}.Index)

		routerWithLogin.POST("/file", controller.FileController{}.Add)
		routerWithLogin.DELETE("/file", controller.FileController{}.Delete)
		routerWithLogin.PUT("/file", controller.FileController{}.Update)
		routerWithLogin.GET("/file", controller.FileController{}.Index)
	}
}
