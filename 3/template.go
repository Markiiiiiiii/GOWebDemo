package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//解析模板
		t, err := template.ParseFiles("./hello.tmpl")
		if err != nil {
			fmt.Println("load template file faild,err:", err)
			return
		}
		// 渲染模板
		name := "Jack"
		err = t.Execute(w, name)
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
