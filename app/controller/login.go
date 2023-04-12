package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"web_standard/app/model"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	var param = make(map[string]string)

	json.NewDecoder(c.Request.Body).Decode(&param)

	fmt.Println(param)

	var user model.User

	//md5加密password
	m := md5.New()
	m.Write([]byte(param["password"]))
	password := hex.EncodeToString(m.Sum(nil))

	con.db().Where("username = ? and password = ?", param["username"], password).First(&user)

	fmt.Println(user)

	if user.Id > 0 {
		//生成token
		var token = user.CreateToken(user.Id)

		//更新token
		con.db().Model(&user).Update("token", token)

		//更新登录时间
		con.db().Model(&user).Update("lasttime_at", time.Now())

		//返回用户数据 以及 token 不返回密码
		user.Password = ""
		con.Success(c, user)
	} else {
		con.Error(c, "用户名或密码错误")
	}
}

func (con LoginController) Add(c *gin.Context) {
	con.Error(c, "test fail")
}
