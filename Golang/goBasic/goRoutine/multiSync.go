/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 09:39:21
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 10:06:02
 * @Description: main
 * @FilePath: /go/src/github.com/liuyuchen777/goRoutine/multiSync.go
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
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
