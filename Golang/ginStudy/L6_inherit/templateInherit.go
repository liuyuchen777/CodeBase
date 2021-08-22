/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 12:28:52
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-30 13:10:52
 * @Description: Say Something, Motherfuncker
 * @FilePath: /Local_Lab/goStudy/gin_demo/templateInherit.go
 * @GitHub: https://github.com/liuyuchen777
 */

// 模板的继承
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexDemo(w http.ResponseWriter, r *http.Request) {
	// 定义

	// 解析
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template error: %v", err)
	}
	msg := "index"
	// 渲染
	t.Execute(w, msg)
}

func homeDemo(w http.ResponseWriter, r *http.Request) {
	// 定义

	// 解析
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template error: %v", err)
	}
	msg := "home"
	// 渲染
	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/home", homeDemo)
	http.HandleFunc("/index", indexDemo)
	fmt.Println("index, click: http://172.28.131.37:9090/index")
	fmt.Println("home, click: http://172.28.131.37:9090/home")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
