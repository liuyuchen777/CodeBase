/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 19:25:18
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 19:34:04
 * @Description: go channel study
 * @FilePath: /go/src/github.com/liuyuchen777/goStudy/goChannel/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // channel is strongly type
	wg.Add(2)
	go func() { // receiver
		i := <-ch
		fmt.Print(i)
		wg.Done()
	}()
	go func() { // sender
		i := 42
		ch <- i
		i = 27
		wg.Done()
	}()
	wg.Wait()
}
