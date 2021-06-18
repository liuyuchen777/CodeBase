/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 09:35:38
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 09:39:12
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/github.com/liuyuchen777/goRoutine/wait.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg string = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "GoodBye"
	wg.Wait() // synchronize two go routine together
}
