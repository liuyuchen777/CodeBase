/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-26 09:12:31
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-03-26 17:20:25
 * @Description: start to learn interface in go
 * @FilePath: /go/src/github.com/liuyuchen777/goInterface/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human // 继承
	// man Human // 类的元素
	score int
	loan  float64
}

type Employee struct {
	Human
	money  float32
	salary float32
}

func (m Human) sayHello() {
	fmt.Println("Hello, I am Human!")
}

func (m Human) sayDetail() {
	fmt.Printf("My name is %v and my age is %v\n", m.name, m.age)
}

func (stu Student) sayHello() {
	fmt.Println("Hello, I am student!")
}

type manInter interface {
	sayHello()
	sayDetail()
}

func main() {
	man := Human{
		name: "Liu Yuchen",
		age:  22,
	}

	var i manInter = man

	man.sayHello()
	i.sayHello()

	stu := Student{Human{"Zhao Zihao", 24}, 97, 10000.7}
	stu.sayHello()
	stu.sayDetail()

	var j manInter = stu
	j.sayDetail()
}
