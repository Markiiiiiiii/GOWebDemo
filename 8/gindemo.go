package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("./xxx", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("./templates/**/*")

	// r.LoadHTMLFiles("./templates/user/index.tmpl") //模板解析
	r.GET("/post/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/index.tmpl", gin.H{ //模板渲染
			"title": "post String",
		})
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{ //模板渲染
			"title": "<a href='http://www.baidu.com'>baidu</a>",
		})
	})

	//返回从网上下载的模板

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.Run(":9000")
}
