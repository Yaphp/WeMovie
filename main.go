package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"web_standard/app/route"
	"web_standard/app/utils"
)

func main() {
	// 设置gin模式
	gin.SetMode(gin.DebugMode)

	//初始化gin实例
	r := gin.Default()

	// 不同模式设置静态文件目录
	if gin.Mode() == "release" {
		utils.Openurl("http://127.0.0.1")
		r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	} else {
		// 服务启动后打开浏览器
		//utils.Openurl("http://localhost:8080")
		r.Use(static.Serve("/", static.LocalFile("./build/dist", false)))
	}

	// 注册上传文件夹路径
	r.Static("/uploads", "./uploads")

	//挂载API路由
	route.ApiRouter(r)

	//输出服务地址
	fmt.Println("Server is running at http://127.0.0.1")

	// tips: 服务启动后打开浏览器
	fmt.Println("如果浏览器没有自动打开，请手动在浏览器中输入http://127.0.0.1访问")

	//启动服务
	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}
