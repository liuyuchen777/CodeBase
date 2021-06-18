/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 12:18:48
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-30 12:30:49
 * @Description: Say Something, Motherfuncker
 * @FilePath: /Local_Lab/goStudy/gin_demo/templateNest.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	// go语言通过首字母是否大小写判断是否对外可用
	Name   string
	Gender string
	Age    int
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/t.tmpl", "./template/ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/tmpl", tmplDemo)
	fmt.Println("start running, click: http://172.28.131.37:9090/tmpl")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
