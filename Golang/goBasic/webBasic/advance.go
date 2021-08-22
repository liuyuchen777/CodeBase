/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 20:55:09
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 21:28:32
 * @Description: form
 * @FilePath: /go/src/github.com/liuyuchen777/goStudy/webBasic/main2.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var myToken string

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// 根路由下的一定会执行？
	/*
		r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		// 注意:如果没有调用 ParseForm 方法，下面无法获取表单的数据
		fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
	*/
	log.Println("take this sayHelloName!")
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix() // 通过当前时间设定md5值
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		myToken = fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(myToken)
		t, _ := template.ParseFiles("/home/lyc/go/src/github.com/liuyuchen777/goStudy/webBasic/login.gtpl")
		t.Execute(w, myToken)
	} else if r.Method == "POST" {
		// 请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证 token 的合法性
			fmt.Println("send file's token is:", token)
			fmt.Println("it should be:", myToken)
			if token == myToken { // 对比token
				fmt.Println("username length:", len(r.Form["username"][0]))
				fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 输出到服务器端
				fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
				fmt.Fprintf(w, "%v login successfully!", r.Form["username"][0])
			}
		} else {
			// 不存在 token 报错
			log.Println("no token exist!")
		}
	} else {
		log.Println("Come to some strange place!")
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	http.HandleFunc("/login", login)   // 设置访问的路由
	fmt.Println("The HP: http://localhost:9090")
	fmt.Println("The login page: http://localhost:9090/login")
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
