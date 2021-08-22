/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-31 09:41:00
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-30 12:23:07
 * @Description: Say Something, Motherfuncker
 * @FilePath: /Local_Lab/goStudy/gin_demo/httpPack.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./template/hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	fmt.Println("start running, click: http://172.28.131.37:9090/hello")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
