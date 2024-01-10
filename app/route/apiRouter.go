package route

import (
	"github.com/gin-gonic/gin"
	"weapp/app/controller"
	"weapp/app/model"
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
		routerWithOutLogin.POST("/register", controller.LoginController{}.Add) // 注册
		routerWithOutLogin.POST("/login", controller.LoginController{}.Index)  // 登录

		routerWithOutLogin.GET("/share", controller.ShareController{}.Index) // 分享文件查询
	}

	routerWithLogin := r.Group("/api").Use(Auth())
	{
		routerWithLogin.POST("/upload", controller.UploadController{}.Index)       // 上传文件
		routerWithLogin.POST("/upload/chunk", controller.UploadController{}.Chunk) // 上传文件

		routerWithLogin.POST("/user", controller.UserController{}.Add)      // 添加用户
		routerWithLogin.DELETE("/user", controller.UserController{}.Delete) // 删除用户
		routerWithLogin.PUT("/user", controller.UserController{}.Update)    // 更新用户
		routerWithLogin.GET("/user", controller.UserController{}.Index)     // 获取用户列表

		routerWithLogin.POST("/file", controller.FileController{}.Add)              // 添加文件
		routerWithLogin.DELETE("/file", controller.FileController{}.Delete)         // 删除文件
		routerWithLogin.PUT("/file", controller.FileController{}.Update)            // 更新文件
		routerWithLogin.GET("/file", controller.FileController{}.Index)             // 获取文件列表
		routerWithLogin.PUT("/file/favorite", controller.FileController{}.Favorite) // 收藏/取消收藏文件

		routerWithLogin.POST("/share", controller.ShareController{}.Add)       // 创建分享文件
		routerWithLogin.POST("/share/save", controller.ShareController{}.Save) // 保存分享文件

	}
}
