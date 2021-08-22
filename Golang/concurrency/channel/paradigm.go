/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 20:54:13
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 20:57:20
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/channel/paradigm.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import "fmt"

func main() {
	// producer
	chanOwner := func() <-chan int { // return a read only channel
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream) // producer take reponsibility to close channel
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // return a closed channel, that consumer only can read from it
	}

	// consumer
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
	// close a read only channel will have compile error
}
