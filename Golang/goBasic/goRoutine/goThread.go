/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 10:09:14
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 10:15:51
 * @Description: Say Something, Motherfuncker
 * @FilePath: /go/src/github.com/liuyuchen777/goRoutine/goThread.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// core of the current machine
	runtime.GOMAXPROCS(1) // set the max thread of this program to 1
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
