/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 19:39:40
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 20:33:24
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/channel/bufferChan.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("---------Channel with Buffer--------")
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Procedure Done!")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received: %v\n", integer)
	}
}
