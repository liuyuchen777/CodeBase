/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 20:32:02
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 20:53:53
 * @Description: basic golang web API
 * @FilePath: /go/src/github.com/liuyuchen777/goStudy/webBasic/main1.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm() //解析参数
	// fmt.Println(r.Form)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("Val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello LYC!")
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello the Next!")
}

func main() {
	http.HandleFunc("/", sayHelloName) // 设置访问的路由
	http.HandleFunc("/lyc", sayHello2)
	fmt.Println("http://localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
