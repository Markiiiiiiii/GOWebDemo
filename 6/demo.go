package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tpl")
	if err != nil {
		fmt.Println("err:,", err)
		return
	}
	msg := "this index page"
	t.Execute(w, msg)
}
func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./home.tpl")
	if err != nil {
		fmt.Println("err:,", err)
		return
	}
	msg := "this home page"
	t.Execute(w, msg)
}
func index2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tpl", "./templates/index.tpl")
	if err != nil {
		fmt.Println("err:,", err)
		return
	}
	msg := "this index page"
	t.ExecuteTemplate(w, "index.tpl", msg)
}
func home2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tpl", "./templates/home.tpl")
	if err != nil {
		fmt.Println("err:,", err)
		return
	}
	msg := "this home page"
	t.ExecuteTemplate(w, "home.tpl", msg)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
