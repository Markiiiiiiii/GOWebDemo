package routers

import (
	"github.com/gin-gonic/gin"
	"gowebdemo.com/demo/21project/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("./static", "static")
	r.LoadHTMLFiles("./templattes/index.html")
	r.GET("/", controller.IndexHandler)
	v1Group := r.Group("./v1")
	{
		//todo
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看
		//1.1查看所有的代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
