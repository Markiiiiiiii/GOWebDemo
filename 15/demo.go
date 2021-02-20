package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//重定向
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.GET("/a", func(c *gin.Context) {
		// 跳转到b对应的路由处理

		c.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(c)        //继续后续在b的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "b",
		})
	})
	r.Run(":9000")
}
