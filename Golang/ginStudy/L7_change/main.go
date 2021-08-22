/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 15:06:34
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-31 15:36:14
 * @Description: Say Something, Motherfuncker
 * @FilePath: /Local_Lab/goStudy/gin_demo/L7_change/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexDemo(w http.ResponseWriter, r *http.Request) {
	// define and parse
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./template/index.tmpl")
	if err != nil {
		fmt.Printf("parse file error: %v", err)
	}
	// rendor
	msg := "刘昱辰"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("rendor error: %v", err)
	}
}

func xssDemo(w http.ResponseWriter, r *http.Request) {
	// define
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./template/xss.tmpl")
	if err != nil {
		fmt.Printf("parse file error: %v", err)
	}
	// rendor
	// tmpl把内容转义了
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='http://liuyuchen777.github.io'>liuyuchenBlog</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Printf("rendor error: %v", err)
	}
}

func xssDemo2(w http.ResponseWriter, r *http.Request) {
	// define
	// parse
	t, err := template.ParseFiles("./template/xss.html")
	if err != nil {
		fmt.Printf("parse file error: %v", err)
	}
	// rendor
	// tmpl把内容转义了
	str := "<a href='liuyuchen777.github.io'>liuyuchenBlog</a>"
	err = t.Execute(w, str)
	if err != nil {
		fmt.Printf("rendor error: %v", err)
	}
}

func main() {
	http.HandleFunc("/index", indexDemo)
	http.HandleFunc("/xss", xssDemo)
	http.HandleFunc("/xss2", xssDemo2)
	fmt.Println("start running, click: http://172.23.174.9:9090/index")
	fmt.Println("start running, click: http://172.23.174.9:9090/xss")
	fmt.Println("start running, click: http://172.23.174.9:9090/xss2")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
