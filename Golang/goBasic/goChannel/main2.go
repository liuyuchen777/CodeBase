/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 19:36:26
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 19:49:24
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/github.com/liuyuchen777/goStudy/goChannel/main2.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // channel is strongly type
	// interior data store
	wg.Add(2)
	go func(ch <-chan int) {
		// receive only channel
		/*
			for i := range ch {
				fmt.Println(i)
			}
		*/
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		// ch <- 27
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		// send only channel
		ch <- 42
		// close(ch) // send message on a closed channel caused panic
		ch <- 27 // block on this line
		ch <- 18
		ch <- 101
		// fmt.Println(<-ch)
		defer close(ch) // close the channel, or will have dead lock on above func
		wg.Done()
	}(ch)
	wg.Wait()
}
