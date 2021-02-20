package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "del",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})
	//noroute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "404 not found",
		})
	})
	//路由组
	vidioGroup := r.Group("/video")
	{
		vidioGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "get",
			})
		})
		vidioGroup.POST("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "post",
			})
		})
		//路由组嵌套
		xx := vidioGroup.Group("/xx")
		xx.POST("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "post",
			})
		})
	}

	r.Run(":9000")
}
