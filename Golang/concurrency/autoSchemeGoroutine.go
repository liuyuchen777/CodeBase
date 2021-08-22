/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-17 08:52:25
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-17 09:00:58
 * @Description:
 * @FilePath: /CodeBase/Golang/concurrency/autoSchemeGoroutine.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"str1", "str2", "str3", "str4"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	/*
		result:
		str4
		str4
		str4
		str4
		goroutine is running a closure that has closed over the iteration variable salutation,
		which has a type of string. As our loop iterates, salutation is being assigned to the
		next string value in the slice literal. Because the goroutines being scheduled may run
		at any point in time in the future, it is undetermined what values will be printed from
		within the goroutine.
		go will autmoatical move salutation to the memory heap to prevent error
	*/

	fmt.Println("---------Refine Version--------")
	for _, salutation := range []string{"str1", "str2", "str3", "str4"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
