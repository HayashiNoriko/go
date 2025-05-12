package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件，生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")

	if err != nil {
		fmt.Println("创建模板失败，err：", err)
		return
	}

	// 利用给定数据渲染模板，并将结构写入 w
	// 1. 传递字符串
	// tmpl.Execute(w, "tina")

	// 2. 传递结构体
	woman := Woman{
		Name: "tina",
		Age:  18,
	}
	tmpl.Execute(w, woman)

}

type Woman struct {
	Name string
	Age  int
}
