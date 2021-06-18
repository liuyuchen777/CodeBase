/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 09:38:07
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 10:18:29
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/github.com/liuyuchen777/goRoutine/anonymous.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	// with anonymous function
	var msg string = "Hello"
	go func() {
		fmt.Println(msg)
	}() // copy
	msg = "GoodBye"
	time.Sleep(100 * time.Millisecond) // not good practice
}

// go run -race xxx.go
// check concurrency