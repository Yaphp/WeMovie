package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"web_standard/app/utils"
)

var Db *gorm.DB
var err error

func init() {
	config := utils.ReadSettingsFromFile("config.json")

	fmt.Println("加载数据库配置信息：" + config.Database.Dsn)

	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Database.Dsn, // DSN data source name
		DefaultStringSize:         256,                 // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,               // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Database.Prefix, // 表名前缀
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("success to connect database ")
}
