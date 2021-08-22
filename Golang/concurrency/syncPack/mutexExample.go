/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 09:17:13
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 09:44:37
 * @Description: example for Mutex, RWMutex in go
 * @FilePath: /CodeBase/Golang/concurrency/mutexExample.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	fmt.Println("---------Mutex-------")
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Increment: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrement: %d\n", count)
	}

	// increment
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// decrement
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Println("-------Mutex Complete------")

	/*
		most of time, mutex lock will degrade system performance
		and most time we don't need to read and write shared variable/memory at the same time
		-> RWMutex
		producer and consumer model
	*/
	fmt.Println("------RWMutex-----")
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMytex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMytex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw, "%d\t%v\t%v\n", count,
			test(count, &m, m.RLocker()), test(count, &m, &m),
		)
	}
	//RWMutex have some advantages on Mutex
}
