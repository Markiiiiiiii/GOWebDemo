package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gowebdemo.com/demo/21project/models"
)

/*
url-->controller-->logic-->model
请求来了-->控制器-->业务逻辑-->模型层的增删改查
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//前端页面填写待办事项，点击提交，发请求到这里
	//从请求中获取数据
	var todo models.Todo
	c.BindJSON(&todo)
	//存入数据库
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "faild",
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	//返回相应
}

func GetTodoList(c *gin.Context) {
	//查询todos表里的所有数据
	// var todolist []Todo
	todolist, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "faild",
		})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "error id"})
		return
	}

	// var todo Todo
	todo, err := models.GetATodoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "faild",
		})
		return
	}

	c.BindJSON(&todo)
	if err := models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "faild",
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	if err := models.DeleteATodoByID(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "faild",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
		})
	}
}
