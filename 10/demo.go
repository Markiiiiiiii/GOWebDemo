package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/query", func(c *gin.Context) {
		//获取浏览器传递的query参数

	})
	r.Run(":900")
}
