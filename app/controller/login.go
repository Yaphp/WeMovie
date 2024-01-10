package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"weapp/app/model"
	"weapp/app/utils"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	var param = make(map[string]string)

	json.NewDecoder(c.Request.Body).Decode(&param)

	//fmt.Println(param)

	var user model.User

	//md5加密password
	m := md5.New()
	m.Write([]byte(param["password"] + "810169"))
	password := hex.EncodeToString(m.Sum(nil))

	model.Db.Where("username = ? and password = ?", param["username"], password).First(&user)

	//fmt.Println(user)

	if user.Id > 0 {
		//生成token
		var token = user.CreateToken(user.Id)

		//更新token
		model.Db.Model(&user).Update("token", token)

		//返回用户数据 以及 token 不返回密码
		user.Password = ""

		// 统计使用的内存总量
		var result []int64
		var total int64
		model.Db.Model(&model.File{}).Where("uid = ?", user.Id).Pluck("Size", &result)
		for _, v := range result {
			total += v
		}
		con.SuccessWithCount(c, user, total)
	} else {
		con.Error(c, "用户名或密码错误")
	}
}

func (con LoginController) Add(c *gin.Context) {
	// 接受参数
	var param = make(map[string]string)

	json.NewDecoder(c.Request.Body).Decode(&param)

	//fmt.Println(param)

	// 判断用户名是否存在
	var user model.User

	model.Db.Where("username = ?", param["username"]).First(&user)

	if user.Id > 0 {
		con.Error(c, "用户名已存在")
		return
	}

	// 注册用户
	//md5加密password
	m := md5.New()
	m.Write([]byte(param["password"] + "810169"))
	password := hex.EncodeToString(m.Sum(nil))

	user.Username = param["username"]
	user.Password = password
	user.CreatedAt = model.LocalTime(utils.GetDateTime())
	user.UpdatedAt = model.LocalTime(utils.GetDateTime())
	user.Role = 1

	model.Db.Create(&user)

	// 判断是否注册成功
	if user.Id > 0 {
		con.Success(c, user)
	} else {
		con.Error(c, "注册失败")
	}

}
