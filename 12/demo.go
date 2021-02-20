package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//获取URI/path参数 E.x:http://127.0.0.0:900/name/age
	r.GET("/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9000")
}
