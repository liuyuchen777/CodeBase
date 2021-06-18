/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 10:43:54
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-30 12:22:47
 * @Description: HTML template
 * @FilePath: /Local_Lab/goStudy/gin_demo/htmlRendor.go
 * @GitHub: https://github.com/liuyuchen777
 */

// 模板与渲染，前后端不分离
// 渲染 -> 字符串替换

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello2(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./template/hello.tmpl")
	if err != nil {
		fmt.Printf("parse file error: %v", err)
		return
	}
	// 渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", sayHello2)
	fmt.Println("start running, click: http://172.28.131.37:9090/hello")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
