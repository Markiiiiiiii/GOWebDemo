package main

import (
	"gowebdemo.com/demo/21project/dao"
	"gowebdemo.com/demo/21project/models"
	"gowebdemo.com/demo/21project/routers"
)

func main() {
	//创建数据库
	//sql: create database todo
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":9090")
}
