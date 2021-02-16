package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//解析模板
		t, err := template.ParseFiles("./person.tmpl")
		if err != nil {
			fmt.Println("load template file faild,err:", err)
			return
		}
		// 渲染模板
		user := Person{
			Name:   "王五",
			Age:    24,
			Gender: "男",
		}
		m1 := map[string]interface{}{
			"Name":   "张三",
			"Age":    25,
			"Gender": "女",
		}
		err = t.Execute(w, map[string]interface{}{
			"user": user,
			"m1":   m1,
		})
		if err != nil {
			fmt.Println("execute faild,err:", err)
			return
		}
	})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server start faild,err", err)
		return
	}
}
