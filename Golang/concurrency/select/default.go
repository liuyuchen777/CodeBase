/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 21:15:21
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 21:20:24
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/select/default.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---------Select Default--------")
	{
		start := time.Now()
		c1 := make(chan interface{})
		c2 := make(chan interface{})
		select {
		case <-c1:
		case <-c2:
		default:
			fmt.Printf("In default after %v.\n", time.Since(start))
		}
	}

	// another example
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
