package model

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int            `json:"id" gorm:"column:id;type:int(11);not null;primaryKey;autoIncrement"`
	Username  string         `json:"username" gorm:"column:username;type:varchar(255);not null"`
	Avatar    string         `json:"avatar" gorm:"column:avatar;type:varchar(255);not null;default:'/images/avatar.png'"`
	Phone     string         `json:"phone" gorm:"column:phone;type:varchar(11);not null"`
	Email     string         `json:"email" gorm:"column:email;type:varchar(50);not null"`
	Password  string         `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Role      int            `json:"role" gorm:"column:role;type:int(11);not null"` // 为0时为超级管理员 为1时为普通管理员 为0时使用Save更新
	Token     string         `json:"token" gorm:"column:token;type:varchar(255);not null"`
	CreatedAt LocalTime      `json:"created_at" gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt LocalTime      `json:"updated_at" gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:datetime;default:null"`
}

// 生成token
func (user User) CreateToken(uid int) string {
	//使用md5 salt+uid+当前时间戳
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%d%d%d", 810169, uid, time.Now().Unix())))
	token := fmt.Sprintf("%x", m.Sum(nil))
	return token
}
