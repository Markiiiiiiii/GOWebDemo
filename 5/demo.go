package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数
	//要么只有一个返回值，要有两个返回值，其中一个必须是error类型
	ku := func(arg string) (string, error) {
		return arg + "so cool", nil
	}

	t := template.New("index.tpl")
	t.Funcs(template.FuncMap{
		"ku99": ku,
	})
	// 在解析模板之前要告诉模板引擎多了一个自定义的函数ku
	_, err := t.ParseFiles("./index.tpl")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	name := "li"
	t.Execute(w, name)
	// 渲染模板
}
func demo1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("t.tpl", "li.tpl")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	name := "ssss"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
