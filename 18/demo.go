package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:123456@(192.168.101.76:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// defer db.
	db.AutoMigrate(&UserInfo{})
	// 添加数据
	// u1 := UserInfo{1, "Mark", "man", "game"}
	// db.Create(u1)

	// 查询数据
	var u UserInfo
	db.First(&u)
	fmt.Println(u)

	// 更新数据
	// db.Model(&u).Update("hobby", "sport")
	// db.First(&u)
	// fmt.Println(u)

	//删除数据
	db.Delete(&u)
}
