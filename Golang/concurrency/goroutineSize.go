/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 09:01:24
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 09:07:14
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/goroutineSize.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000) // calculate memory that 1 goroutine take
}
