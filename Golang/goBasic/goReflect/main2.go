/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 17:21:31
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 17:23:37
 * @Description: test of interface in more practical
 * @FilePath: /go/src/github.com/liuyuchen777/goInterface/main2.go
 * @GitHub: https://github.com/liuyuchen777
 */

/*
prototype defination of string() in fmt.Println

type Stringer interface {
    String() string
}
*/

package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
