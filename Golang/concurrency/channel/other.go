/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 20:42:28
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 20:53:01
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/channel/other.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import "fmt"

func main() {
	fmt.Println("-----------Some Test------------")
	var c chan int
	if c == nil {
		fmt.Println("Now we have a nil channel!")
	}
	c = make(chan int)
	if c != nil {
		fmt.Println("Now we have a non-nil channel!")
	}
	close(c)
	// can't reopen a closed channel
	stringStream := make(chan string, 4)
	stringStream <- "LYC"
	stringStream <- "WZH"
	close(stringStream)
	fmt.Println(<-stringStream)
	fmt.Println(<-stringStream)
	fmt.Println(<-stringStream)
}
