package mysql

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlInit(conf *ini.File) {
	user := conf.Section("mysql").Key("user").String()
	password := conf.Section("mysql").Key("password").String()
	ip := conf.Section("mysql").Key("ip").String()
	port, _ := conf.Section("mysql").Key("port").Int()
	database := conf.Section("mysql").Key("database").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		//SkipDefaultTransaction: true, //禁用事务
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
}
