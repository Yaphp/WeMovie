package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"wemovie/app/route"
	"wemovie/app/utils"
)

func main() {
	// 设置gin模式
	gin.SetMode(gin.ReleaseMode)

	//初始化gin实例
	r := gin.Default()

	// 不同模式设置静态文件目录
	if gin.Mode() == "release" {
		utils.Openurl("https://127.0.0.1")
		r.Use(static.Serve("/", static.LocalFile(utils.GetRootPath()+"/dist", false)))
	} else {
		// 服务启动后打开浏览器
		//utils.Openurl("http://localhost:8080")
		r.Use(static.Serve("/", static.LocalFile(utils.GetRootPath()+"/dist", false)))
	}

	//挂载API路由
	route.ApiRouter(r)

	// tips: 服务启动后打开浏览器
	fmt.Println("如果浏览器没有自动打开，请手动在浏览器中输入https://127.0.0.1访问")

	//启动服务
	r.RunTLS(":443", utils.GetRootPath()+"/server.crt", utils.GetRootPath()+"/server.key") // 监听并在 https://0.0.0.0:443 上启动服务
}
