package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID int64
	// Name string `gorm:"default:'iii'"` //指定一个默认值，不传值或者0值都会使用默认值
	Name *string `gorm:"default:'iii'"` //指定一个默认值，不传值或者0值时可以保存为空值或0值
	Age  int64
}

func main() {
	dsn := "root:123456@(192.168.101.76:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//生成表
	// db.AutoMigrate(&User{})
	//新增一个结构体的实例
	// name := "time"
	// u := User{Name: &name, Age: 20}
	// db.Create(&u)

	//查询
	// var u User //创建一个变量承接返回的数据
	// db.First(&u)
	// fmt.Println(u)

	// var us []User
	// db.Find(&us)
	// fmt.Println(us)

	// db.Debug().First(&us, 3) //指定查询第3条
	// fmt.Println(us)

	// db.Where("age=?", 20).First(&us)
	// fmt.Println(us)

	var us User
	name := "xxxx"
	db.Debug().FirstOrCreate(&us, User{Name: &name})
	fmt.Printf("%#v\n", us)
}
