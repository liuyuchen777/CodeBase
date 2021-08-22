/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 18:35:42
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 18:45:00
 * @Description: example for once
 * @FilePath: /CodeBase/Golang/concurrency/onceExample.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		sync.Once assure that even on different goroutines,function execute one time
	*/
	var count int
	increment := func() {
		count++
	}
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()
	fmt.Print("Count: ", count)

	/*
		once only calculate times of do for once no matter what function
	*/
	var count2 int
	increments := func() { count2++ }
	decrements := func() { count2-- }
	var onces sync.Once
	onces.Do(increments)
	onces.Do(decrements)
	fmt.Printf("Count2: %d\n", count2)

	// an example of deadlock
	{
		var onceA, onceB sync.Once
		var initB, initA func()
		initA = func() { onceB.Do(initB) }
		initB = func() { onceA.Do(initA) }
		onceA.Do(initA)
	}
}
