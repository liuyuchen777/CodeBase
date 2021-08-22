/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 08:49:49
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 08:52:02
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/copyInGoroutine.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "Welcome"
	}()
	wg.Wait()
	fmt.Print(salutation)
	/*
		you will see welcome
		it shows that goroutine is run in the same address space they were created in
	*/
}
