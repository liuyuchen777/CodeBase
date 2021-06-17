/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 08:39:46
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 08:45:41
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/simpleJoinTest.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World!")
}

func main() {
	go sayHello()
	time.Sleep(2 * time.Second)
}
