/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 11:22:32
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-30 12:17:31
 * @Description: Say Something, Motherfuncker
 * @FilePath: /Local_Lab/goStudy/gin_demo/htmlRendor2.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfor struct {
	// go语言通过首字母是否大小写判断是否对外可用
	Name string
	Age  int
	Say  string
}

func sayHello3(w http.ResponseWriter, r *http.Request) {
	// self define function
	cool := func(name string) (string, error) {
		return name + " is so cool", nil
	}
	// 定义模板
	t := template.New("./template/hello2.tmpl")
	// register func，一定在解析模板之前
	t.Funcs(template.FuncMap{
		"myFun": cool,
	})
	// 解析模板
	_, err := t.ParseFiles("./hello2.tmpl")
	if err != nil {
		fmt.Printf("parse file error: %v", err)
		return
	}
	// 渲染模板
	// 可以传结构体
	u1 := UserInfor{
		Name: "lyc",
		Age:  22,
		Say:  "I am boss",
	}
	// 也可以传map
	m1 := map[string]interface{}{
		"Name": "lyc",
		"Age":  22,
		"Say":  "I am King",
	}
	err = t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"u2":    m1,
		"hobby": []string{"football", "run", "swim", "walk"},
	})
	if err != nil {
		fmt.Printf("render error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", sayHello3)
	fmt.Println("start running, click: http://172.28.131.37:9090/hello")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
