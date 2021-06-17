/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 18:48:44
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 19:16:49
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/syncPack/poolExample.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		object pool model
	*/
	fmt.Println("---------Simple Pool Example-------")
	{
		myPool := &sync.Pool{
			New: func() interface{} {
				fmt.Println("Creating new instance.")
				return struct{}{}
			},
		}

		myPool.Get()
		instance := myPool.Get()
		myPool.Put(instance)
		myPool.Get()

		// only create two instance
	}

	fmt.Println("----------Useful Case----------")
	{
		// why??????
		var numCalcsCreated int

		calcPool := &sync.Pool{
			New: func() interface{} {
				numCalcsCreated += 1
				mem := make([]byte, 1024)
				return &mem // 1
			},
		}
		// create 4K memory
		calcPool.Put(calcPool.New())
		calcPool.Put(calcPool.New())
		calcPool.Put(calcPool.New())
		calcPool.Put(calcPool.New())
		const numWorkers = 1024 * 1024
		var wg sync.WaitGroup
		wg.Add(numWorkers)
		for i := numWorkers; i > 0; i-- {
			go func() {
				defer wg.Done()
				mem := calcPool.Get()
				defer calcPool.Put(mem)

				// assume something quick is being done with this memory
			}()
		}

		wg.Wait()
		fmt.Printf("%d calculators were created.", numCalcsCreated)

		// every execution, 8 calculators were created
	}
}
