/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 20:57:48
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 21:12:27
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/select/basicSelect.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("-------------Basic Example-----------")
	// dead loop example
	{
		start := time.Now()
		c1 := make(chan interface{})
		c2 := make(chan interface{})

		go func() {
			for {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				c1 <- struct{}{}
			}
		}()

		go func() {
			for {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				c2 <- struct{}{}
			}
		}()

		fmt.Println("Blocking on read...")
		for {
			select {
			case <-c1: // 2
				fmt.Printf("[f1] Unblocked %v later.\n", time.Since(start))
			case <-c2:
				fmt.Printf("[f2] Unblocked %v later.\n", time.Since(start))
			}
		}
	}
}
