/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 09:24:32
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 09:37:58
 * @Description: start to learn go routine
 * @FilePath: /go/src/github.com/liuyuchen777/goRoutine/simple.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"time"
)

// creating go routine

// synchronization: 1. wait groups, 2. mutexes

// parallelism

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
