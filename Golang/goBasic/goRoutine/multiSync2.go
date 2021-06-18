/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 10:06:00
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 10:06:38
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/github.com/liuyuchen777/goRoutine/multiSync2.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0 // shared object
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
