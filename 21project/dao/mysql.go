package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB //全局数据库连接对象
)

func InitMysql() (err error) {
	DB, err = gorm.Open(mysql.Open("root:qaz78963@(192.168.199.125:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("MySQL content faild! err:", err)
		return
	}
	return
}
