package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/query", func(c *gin.Context) {
		//获取浏览器传递的query参数
		// name := c.Query("query")
		// name := c.DefaultQuery("query", "somebady")
		name, ok := c.GetQuery("query")
		if !ok {
			name = "somebady"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9000")
}
