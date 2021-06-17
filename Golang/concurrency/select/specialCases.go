/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 21:12:51
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 21:20:54
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/select/specialCases.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----------At Same Time--------")
	// what if several channel needs to be read at same time -> random choice by go
	{
		c1 := make(chan interface{})
		close(c1)
		c2 := make(chan interface{})
		close(c2)
		var c1Count, c2Count int
		for i := 1000; i >= 0; i-- {
			select {
			case <-c1:
				c1Count++
			case <-c2:
				c2Count++
			}
		}
		fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
	}

	fmt.Println("------------No Channel is Initialized---------")
	{
		var c <-chan int
		select {
		case <-c: //1
		case <-time.After(1 * time.Second):
			fmt.Println("Timed out.")
		}
	}

	// select without case will block forever
	select {}
}
