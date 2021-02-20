package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.htm")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.htm", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// username := c.PostForm("username")
		// password := c.PostForm("password")
		// username := c.DefaultPostForm("username", "name")
		// password := c.DefaultPostForm("password", "psw")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "abc"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "abc"
		}
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	r.Run(":9000")
}
