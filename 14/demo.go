package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//r最大接受8<<20就是8M
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			// filepath := fmt.Sprintf("./%s", file.Filename)
			filepath := path.Join("./", file.Filename)
			c.SaveUploadedFile(file, filepath)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})
	r.Run(":9000")
}
