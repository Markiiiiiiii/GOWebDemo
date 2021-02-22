package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB //全局数据库连接对象
)

// todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() (err error) {
	DB, err = gorm.Open(mysql.Open("root:qaz78963@(192.168.199.125:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("MySQL content faild! err:", err)
		return
	}
	return
}

func main() {
	//创建数据库
	//sql: create database todo
	//连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("./static", "static")
	r.LoadHTMLFiles("./templattes/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := r.Group("./v1")
	{
		//todo
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项，点击提交，发请求到这里
			//从请求中获取数据
			var todo Todo
			c.BindJSON(&todo)
			//存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "faild",
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
			//返回相应
		})
		// 查看
		//1.1查看所有的代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todos表里的所有数据
			var todolist []Todo
			if err = DB.Find(&todolist).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "faild",
				})
			} else {
				c.JSON(http.StatusOK, todolist)
			}
		})
		//1.2查看所有的代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "faild",
				})
				return
			}

			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "faild",
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "faild",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
				})
			}
		})
	}
	r.Run(":9090")
}
