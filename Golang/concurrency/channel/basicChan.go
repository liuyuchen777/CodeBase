/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 19:18:34
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 20:59:39
 * @Description: basic usage of go channel
 * @FilePath: /CodeBase/Golang/concurrency/channel/basicChan.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		// default channel
		var dataStream chan interface{}
		dataStream = make(chan interface{})

		// single direction channel
		var dataStream <-chan interface{}
		dataStream := make(<-chan interface{})

		var dataStream chan<- interface{}
		datastream = make(chan<- interface{})

		// divide a full channel to two single direction channel
		var receiveChan <-chan interface{}
		var sendChan chan<- interface{}
		dataStream := make(chan interface{})
		receiveChan = dataStream
		sendChan = dataStream
	*/
	fmt.Println("----------Basic Example---------")
	{
		stringStream := make(chan string)
		go func() {
			stringStream <- "Hello Channels"
		}()
		fmt.Println(<-stringStream)
	}

	// the second bool return value from channel shows the channel is closed or not
	stream := make(chan string)
	go func() {
		stream <- "Hello Channels!"
	}()
	receive, ok := <-stream
	fmt.Printf("(%v): %v\n", ok, receive)
	close(stream) // close a channel
	receive, ok = <-stream
	fmt.Printf("(%v): %v\n", ok, receive)

	// common paradigm when use channel
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	for i := range intStream {
		fmt.Println(i)
	}

	// use close to fix block
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // block here
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin) // close channel to send signal
	wg.Wait()
}
