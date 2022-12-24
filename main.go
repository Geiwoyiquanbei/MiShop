package main

import (
	"MiShop/dao/mysql"
	"MiShop/dao/redis"
	"MiShop/logger"
	"MiShop/routers"
	"MiShop/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化设置
	conf, err := settings.Init()
	if err != nil {
		return
	}
	//配置logger
	logger.InitLogger()
	//配置mysql
	mysql.MysqlInit(conf)
	//配置redis
	redis.RedisInit(conf)
	//加载路由
	r := gin.Default()
	routers.SetUp(r)
	err = r.Run()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	return
}
