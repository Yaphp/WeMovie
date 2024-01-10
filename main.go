package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"weapp/app/route"
	"weapp/app/utils"
)

func main() {
	// 设置gin模式
	gin.SetMode(gin.ReleaseMode)

	//初始化gin实例
	r := gin.Default()

	// 设置静态文件目录
	r.Use(static.Serve("/", static.LocalFile(utils.GetRootPath()+"/dist", false)))
	r.Use(static.Serve("/uploads", static.LocalFile(utils.GetRootPath()+"/uploads", false)))

	//挂载API路由
	route.ApiRouter(r)

	// tips: 服务启动后打开浏览器
	fmt.Println("如果浏览器没有自动打开，请手动在浏览器中输入https://127.0.0.1访问")

	config := utils.ReadSettingsFromFile(utils.GetRootPath() + "/config.json")

	// 监听并在 https://0.0.0.0:443 上启动服务
	r.RunTLS(":"+config.App.Port, utils.GetRootPath()+"/server.crt", utils.GetRootPath()+"/server.key")
}
