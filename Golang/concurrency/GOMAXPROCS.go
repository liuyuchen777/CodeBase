/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 21:21:34
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 21:25:28
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/GOMAXPROCS.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// set go to use all CPU in machine
	runtime.GOMAXPROCS(runtime.NumCPU())
	// check current
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}
