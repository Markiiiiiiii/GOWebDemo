package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` //tag定义数据库表中列的属性
	Role         string  `gorm:"size:255"`                       //设置字段大小
	MemberNumber *string `gorm:"unique;not null"`                //设置会员号是唯一并且非空
	Num          int     `gorm:"AUTO_INCREMENT"`                 //设置num为自增类型
	Address      string  `gorm:"index:addr"`                     //为address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                              //忽略本字段
}

func main() {
	dsn := "root:123456@(192.168.101.76:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// defer db.

	db.AutoMigrate(&User{})

	// db.Model(&u).Update("CreatAt",time.Now()) //更新creatat的时间
	// 添加数据
	// u1 := UserInfo{1, "Mark", "man", "game"}
	// db.Create(u1)

	// 查询数据
	// var u UserInfo
	// db.First(&u)
	// fmt.Println(u)

	// // 更新数据
	// // db.Model(&u).Update("hobby", "sport")
	// // db.First(&u)
	// // fmt.Println(u)

	// //删除数据
	// db.Delete(&u)
}
