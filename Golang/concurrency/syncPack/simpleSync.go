/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 08:46:24
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 09:15:15
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/simpleSync.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

func syncSayHello() {
	var wg sync.WaitGroup
	wg.Add(1)
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello World!")
	}
	go sayHello()
	wg.Wait()
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print("i = ", i+1, ": ")
		syncSayHello()
	}

	fmt.Println("----------Another Example---------")
	var wg sync.WaitGroup
	wg.Add(1) //1
	go func() {
		defer wg.Done() //2
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()
	wg.Add(1) //1
	go func() {
		defer wg.Done() //2
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()
	wg.Wait() //3
	fmt.Println("All goroutines complete.")
}
