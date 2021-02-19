package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("index.tpl").Delims("{[", "]}").ParseFiles("./index.tpl")
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		name := "LLL"
		t.Execute(w, name)
	})
	// http.HandleFunc("/xcc", func(w http.ResponseWriter, r *http.Request) {
	// 	t, err := template.New("index.tpl").Funcs(template.FuncMap{
	// 		"safe": func(s string) template.HTML {
	// 			return template.HTML(s)
	// 		},
	// 	}).ParseFiles("./index.tpl")
	// 	if err != nil {
	// 		fmt.Println("err:", err)
	// 		return
	// 	}
	// 	name := "LLL"
	// 	t.Execute(w, name)
	// })
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
