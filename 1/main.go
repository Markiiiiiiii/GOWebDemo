package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //返回默认的路由引擎
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello golang",
		})
	})
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	r.Run()
}
