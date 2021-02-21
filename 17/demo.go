package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandle(c *gin.Context) {
	fmt.Println("index in ...")
	name, ok := c.Get("name") //跨中间件取值
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": name,
		})
	} else {
		fmt.Println("error")
	}
}
func shop(c *gin.Context) {
	fmt.Println("index in ...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}
func user(c *gin.Context) {
	fmt.Println("index in ...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//中间件m1,统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	//  go funcxx(c.Copy())//在中间件中使用gorouting时只能使用c的copy
	c.Next() //调用后续的处理函数，这里开始处理indexhandle
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
	// c.Abort() //阻止调用后续的处理函数

}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "fff") //设置中间件中传递的变量
	c.Next()             //调用后续的处理函数，这里开始处理indexhandle
	// c.Abort() //阻止调用后续的处理函数，这里开始处理indexhandle
	fmt.Println("m2 out ...")

}

//重要！！！！！ 中间件 生产环境中的常用写法
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//其他的一些工作
	return func(c *gin.Context) {
		if doCheck {
			//存放一些具体的逻辑
			//是否判断登陆
		} else {
			c.Next()
		}

	}
}
func testMiddle(c *gin.Context) {
	fmt.Println("m3 in...")
	c.Next()
	fmt.Println("m3 out...")
}

func main() {
	r := gin.Default() //默认使用logger和recovery中间件
	// r:=gin.New()不想使用就直接使用new函数启动一个不包含上面两个中间件的空函数

	// r.Use(m1, m2, authMiddleware(true)) //全局注册中间件函数m1
	r.Use(m1, m2)
	r.GET("/index", indexHandle)
	r.GET("/shop", shop)
	r.GET("/user", user)
	r.GET("/m3", testMiddle, user) //给路由单独注册一个
	//路由组注册方法1
	xx := r.Group("/xx", testMiddle)
	{
		xx.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "ok!"}) })
	}
	//路由组注册方法2
	xxx := r.Group("/xxx")
	xxx.Use(testMiddle)
	{
		xxx.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "ok!"}) })
	}

	r.Run(":9000")
}
