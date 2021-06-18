/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 19:50:00
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 20:25:54
 * @Description: logger
 * @FilePath: /go/src/github.com/liuyuchen777/goStudy/goChannel/main3.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var wg = sync.WaitGroup{}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel

func main() {
	wg.Add(2)
	go logger()
	// time.Sleep(100 * time.Millisecond)
	go sender()
	wg.Wait()
}

func sender() {
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logWarning, "App is shutting down"}
	// time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
	wg.Done()
}

func logger() {
	for {
		select { // if no message in channel, block here
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			wg.Done()
		}
	}
}
