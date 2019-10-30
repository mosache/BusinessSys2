package main

import (
	"BusinessSys/config"
	"BusinessSys/db"
	"BusinessSys/ginengine"
	"BusinessSys/middleware"
	routes "BusinessSys/routers"
	"fmt"
)

func main() {

	//初始化配置
	config.InitConfig()

	//初始化GinEngine
	ginengine.InitEngine()

	//初始化中间件
	middleware.SetUp()

	//初始化routes
	routes.Setup()

	//初始化数据库
	db.InitDB()

	//gin run
	err := ginengine.Engine.Run(fmt.Sprintf(":%d", config.AppConfig.Port))

	if err != nil {
		panic("start gin failed!")
	}
}
