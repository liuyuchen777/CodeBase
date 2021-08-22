/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 09:44:57
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 18:35:33
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/condExample.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		cond is when condition statisfy, execute
	*/
	fmt.Println("-------Signal-------")
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from queue")
		c.L.Unlock()
		c.Signal() // send signal to wait goroutine that you can go to next step
	}

	for i := 0; i < 10; i++ {
		// add something in queue
		c.L.Lock()
		for len(queue) == 2 { // when there are two elements in queue, block the progress
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{}) // queue for empty struct{}
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}

	fmt.Println("Len of queue: ", len(queue))

	/*
		There are two ways to send signal: Signal and broadcast
		signal send signal to the longest waiting goroutine in FIFO list
		broadcast send signal to all the waiting goroutines
	*/
	// broadcast example
	fmt.Println("---------Broadcast--------")
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var tempwg sync.WaitGroup
		tempwg.Add(1)
		go func() {
			tempwg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		tempwg.Wait()
		// c.L.Lock()
		// defer c.L.Unlock()
		// fn()
	}

	var wg sync.WaitGroup
	wg.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing Window!")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Display annoying dialog box!")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked!")
		wg.Done()
	})

	button.Clicked.Broadcast() // button is clicked, broadcast message to all subscriber

	wg.Wait()

	// another way to write below program
	fmt.Println("--------Another Way--------")
	subscribe2 := func(c *sync.Cond, fn func()) {
		c.L.Lock()
		defer c.L.Unlock()
		fn()
	}
	wg.Add(3)
	go subscribe2(button.Clicked, func() {
		fmt.Println("Hello LYC!")
		wg.Done()
	})
	go subscribe2(button.Clicked, func() {
		fmt.Println("Hello ZZH!")
		wg.Done()
	})
	go subscribe2(button.Clicked, func() {
		fmt.Println("Hello XXL!")
		wg.Done()
	})
	button.Clicked.Broadcast()
	wg.Wait()
}
