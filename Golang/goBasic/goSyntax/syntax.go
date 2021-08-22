/*
 * @Author: Liu Yuchen
 * @Date: 2021-03-20 09:31:58
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-18 10:07:09
 * @Description: go play ground
 * @FilePath: /CodeBase/Golang/goBasic/goSyntax/main.go
 * @GitHub: https://github.com/liuyuchen777
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

// global variable need to be full declare
var i float32 = 100
var (
	actorName string = "Liu Yuchen"
	company   string = "Tokyo Tech"
	idNumber  int    = 20
)
var (
	counter int = 10
)

// j := 100 // not able to do that

func varShow() {
	// three method to create variable in go
	fmt.Println(i)
	var i int // shadowing
	i = 42

	var j int = 43 // local variable in go always have to be used

	k := 44

	var l float32 = 27
	// print something
	fmt.Println("Hello World!")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	//printf of go
	fmt.Printf("%v, %T\n", l, l)
	// use function from other package

	// convert variable type
	var intVar int = 42
	fmt.Printf("%v, %T\n", intVar, intVar)
	var jVar float32
	jVar = float32(intVar)
	fmt.Printf("%v, %T\n", jVar, jVar)

	var kVar float32 = 42.5
	intVar = int(kVar)
	// intVar = kVar // can't explicit to do the convert, may loss information
	fmt.Printf("%v, %T\n", intVar, intVar)
}

func primeShow() { // show some primitive data type
	// boolean type
	var t bool = true
	fmt.Println(t)
	m := 1 == 1
	n := 1 == 2
	fmt.Println(m, n)
	fmt.Printf("%v, %T", m, m)
	// int type
	var n1 uint16 = 42 // unsigned int
	var n2 uint32 = 43
	var n3 uint64 = 44
	fmt.Println(n1, n2, n3)
	// define of float
	fl := 3.14
	fl = 13.7e72
	fl = 2.1e14
	fmt.Printf("%v, %T\n", fl, fl)
	// float var operation
	a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	// complex number
	var c1 complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c1, c1)
	var c2 complex64 = 2 + 5.2i
	fmt.Println(c1 + c2)
	fmt.Println(c1 - c2)
	fmt.Println(c1 * c2)
	fmt.Println(c1 / c2)
	// string
	s1 := "I am Liu Yuchen!"
	// s1[2]= "u"
	fmt.Printf("%v, %T\n", string(s1[2]), s1[2])
	bytes := []byte(s1)                        // convert string to slice
	fmt.Printf("%v, %T\n", bytes[0], bytes[0]) //73, uint8
	// r := 'a'
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // 97, int32
	// rune -> UTF32(int32), byte -> UTF8(uint8), alias for int32
	// special function to process
}

const aCon int = 27

const (
	aa = iota // iota is a enumeration type
	bb = iota
	cc = iota
	/* also can do in the following form
	 * aa = iota
	 * bb
	 * cc
	 */
)

const (
	_ = iota // from zero value
	catSpecial
	dogSpecial
	snakeSpecial
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 左移10位
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func constantShow() {
	// declare a constant type
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
	// const myConst2 float64 = math.Sin(1.67) // value of constant need to be determine
	const aCon int = 26
	fmt.Println(aCon)
	// var bVar int16 = 17
	// fmt.Println(aCon + bVar) // int cannot add with int16
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	// enume
	fileSize := 4000000000.
	fmt.Printf("%.2f GB\n", fileSize/GB)
}

func arrayShow() {
	// array: 1. creation 2. built-in function 3. work with arrays
	grades := [3]int{97, 85, 93} // array continuous in memory
	// grades := [...]int{97, 85, 93} // another way to declare array
	var students [3]string
	students[0] = "Lisa"
	students[1] = "Lyc"
	// print
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	students[2] = "Bob"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("studnets #1: %v\n", students[1])
	fmt.Printf("Numbers of Students: %v\n", len(students))
	// 2-D array
	// var Matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	var Matrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	Matrix2 := [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	var Matrix3 [3][3]int
	Matrix3[0] = [3]int{1, 0, 0}
	Matrix3[1] = [3]int{0, 1, 0}
	Matrix3[2] = [3]int{0, 0, 1}
	fmt.Println(Matrix)
	fmt.Println(Matrix2)
	fmt.Println(Matrix3)
	studentsCopy := students
	students[2] = "zzh"
	fmt.Printf("studentsCopy: %v\n", studentsCopy)
	fmt.Printf("students: %v\n", students)
	studentsAlias := &students
	students[0] = "xxl"
	fmt.Printf("studentsAlias: %v\n", studentsAlias)
}

func sliceShow() {
	// slice: 1. creation 2. built-in function 3. work with slice
	a := []int{1, 2, 3} // don't state the size, array will be slice
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	// other way to create slice
	aSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bSlice := aSlice[:]  // slice of all elemnts
	cSlice := aSlice[3:] // slice from 4th to end
	dSlice := aSlice[:6] // slice of first 6 elemnts
	eSlice := aSlice[3:6]
	aSlice[5] = 42 // all change, pointer, not like array
	fmt.Println(aSlice)
	fmt.Println(bSlice)
	fmt.Println(cSlice)
	fmt.Println(dSlice)
	fmt.Println(eSlice)
	// make function to create slice
	aMake := make([]int, 3, 100) // make(sliceType, size, capacity)
	fmt.Println(aMake)
	fmt.Printf("Length: %v\n", len(aMake))
	fmt.Printf("Capacity: %v\n", cap(aMake))
	// append
	aMake = append(aMake, 1)
	fmt.Println(aMake)

	aa := []int{}
	fmt.Println(aa)
	fmt.Printf("Length: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))
	aa = append(aa, 1)
	fmt.Println(aa)
	fmt.Printf("Length: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))
	aa = append(aa, 2, 3, 4, 5)
	fmt.Println(aa)
	fmt.Printf("Length: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))
	aa = append(aa, []int{6, 7, 8, 9}...) // ... seperate slice to several elements
	fmt.Println(aa)
	fmt.Printf("Length: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))
}

func sliceShow2() {
	// method of slice
	//len, cap, append
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...) // b reference first two elements of a
	fmt.Println(b)
	fmt.Println(a)
}

func mapShow() {
	// what? creating? manipulation?
	statePopulation := map[string]int{ // key: value
		"Cal": 100,
		"Tex": 200,
		"Flo": 300,
		"NYC": 400,
		"Pen": 500,
		"Ohi": 600,
	}
	fmt.Println(statePopulation)
	m := map[[3]int]string{
		{1, 2, 3}: "lyc",
	}
	fmt.Println(m)
	// use make function to create map
	countryPopulation := make(map[string]int)
	fmt.Println(countryPopulation)
	//use map
	fmt.Println(statePopulation["NYC"])
	statePopulation["Geo"] = 10000
	fmt.Println(statePopulation["Geo"])
	fmt.Println(statePopulation["LA"]) // no "LA in the map
	delete(statePopulation, "Geo")
	fmt.Println(statePopulation["Geo"]) // no longer "Geo" in the map
	pop, ok := statePopulation["Geo"]
	fmt.Println(pop, ok)
	// how many elements
	fmt.Println(len(statePopulation))
	sp := statePopulation
	delete(sp, "NYC") // not copy, directly reference
	fmt.Println(len(sp))
	fmt.Println(len(statePopulation))
}

type Doctor struct {
	number      int
	actorName   string
	compainions []string // slice
}

// object-oriented feature of Go
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	speedKPH float32
	canFly   bool
}

func structShow() {
	// in struct, you can put any type of data
	aDoctor := Doctor{
		number:    3,
		actorName: "Liu Yuchen",
		compainions: []string{
			"Jiang Ning",
			"Wang Zihan",
			"Wang Yahui",
		},
	}
	/* work but not recommend
	aDoctor := Doctor{
		3,
		"Liu Yuchen",
		[]string{
			"JiangNing",
			"WangZihan",
			"WangYahui",
		},
	}
	*/
	fmt.Println(aDoctor)
	// check element in struct
	fmt.Println(aDoctor.actorName)
	// 匿名struct
	bDoctor := struct{ name string }{name: "Spike Liu"}
	fmt.Println(bDoctor)

	// father and son
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Austrlia"
	b.speedKPH = 48.
	b.canFly = false
	fmt.Println(b.Name)

	b2 := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		speedKPH: 48.,
		canFly:   false,
	}
	fmt.Println(b2.Name, b2.speedKPH, b2.canFly)
	// struct is collections of disparate data types that describe a single concept
	// keyed by named fields
	// anonymous structs are allowed
	// structs are value types
	// no inheritance, but can be embedded
	// tags can be added to struct fields to decribe field
}

func flowControl() {
	// as simple as we can
	if true {
		fmt.Println("the test is true!")
	}
	// common form
	statePopulation := map[string]int{ // key: value
		"Cal": 100,
		"Tex": 200,
		"Flo": 300,
		"NYC": 400,
		"Pen": 500,
		"Ohi": 600,
	}
	if pop, ok := statePopulation["NYC"]; ok {
		// pop and ok will be if local
		fmt.Println(pop)
	}
	// fmt.Println(pop) // error
	// be careful with float numbers
	myNum := 0.123 // 0.1 will be given same value
	// if myNum == math.Pow(math.Sqrt(myNum), 2) {
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("They are same numbers!")
	} else {
		fmt.Println("They are different numbers!")
	}
}

func switchControl() {
	// just like C/C++ switch
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
	// another form
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Print("two, four or six")
	default:
		fmt.Println("another number")
	}
	// another form
	i := 10
	switch {
	case i <= 10:
		fmt.Println("i case one!")
	case i <= 20:
		fmt.Println("i case two!")
	default:
		fmt.Println("i default case!")
	}
	// golang don't need break for switch
	j := 10
	switch {
	case j <= 10:
		fmt.Println("j case one!")
		fallthrough
	case j <= 20:
		fmt.Println("j case two!")
	default:
		fmt.Println("j default case!")
	}
	// fallthrough, not very common used
	// type switch, very special
	var ii interface{} = 1. // interface auto judge type
	switch ii.(type) {
	case int:
		fmt.Println("int!")
		// break
		// fmt.Println("This will not be printed")
	case float64:
		fmt.Println("float64!")
	case string:
		fmt.Println("string!")
	default:
		fmt.Println("i is another type!")
	}
}

func loopShow() {
	// basic loop pattern
	// 1. for define: judge; incremental {}
	for i := 1; i < 3; i++ {
		fmt.Println("【Type 1】 ", i)
	}
	// 2. for judge {}
	j := 1
	for j < 3 {
		fmt.Println("【Type 2】", j)
		j++
	}
	// 3. dead loop
	i := 1
	for {
		fmt.Println("【Type 3】", i)
		i++
		if i == 3 {
			break
		}
	}
	// label define
Loop:
	for i3 := 1; i3 <= 3; i3++ {
		for j3 := 1; j3 <= 3; j3++ {
			fmt.Println("【Label】", i3*j3)
			if i3*j3 >= 3 {
				break Loop
			}
		}
	}
	// for in set: array, slice, string and map
	arr := [3]int{1, 2, 3}
	for k := range arr {
		fmt.Println("【array collection】", k)
	}
	sl := []int{1, 2, 3}
	for k := range sl {
		fmt.Println("【slice collection】", k)
	}
	mp := map[string]int{ // key: value
		"Cal": 100,
		"Tex": 200,
		"Flo": 300,
		"NYC": 400,
		"Pen": 500,
		"Ohi": 600,
	}
	for k, v := range mp {
		fmt.Println("【map collection】", k, v)
	}
}

func defer1() {
	// defer: associate with close resource (for example: in Internet program)
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	// defer is FILO, like stack
}

func defer2() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func defer3() {
	a := "start"
	// defer a line of code, save the variable value at that time
	defer fmt.Println(a)
	a = "end"
}

func panic1() {
	/* some bad error happened
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
	*/
	fmt.Println("start")
	panic("something bad happened!")
	fmt.Println("end")
	// panic make the program can not be continued!
}

func panic2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// throw some error
		panic(err.Error())
	}
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			// recover only useful in defer functions
			// recover is used to revocer from panic
			// current function will not attemptto continue, but highrt function can still continue
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func recover1() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

type myStruct struct {
	foo int
}

func pointerShow() {
	// create pointer
	a := 42
	b := 42 // copy value from a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
	// use pointer!
	var aa int = 42
	var bb *int = &aa
	fmt.Println(aa, *bb)
	aa = 27
	fmt.Println(aa, *bb)
	*bb = 14
	fmt.Println(aa, *bb)
	// array
	a1 := [3]int{1, 2, 3}
	b1 := &a1[0]
	b2 := &a1[1] // address of a1 is constant
	fmt.Printf("%v, %p, %p\n", a1, b1, b2)
	// "unsafe" package
	// struct example
	var mst myStruct = myStruct{foo: 42}
	fmt.Println(mst) // {42}
	mst.foo = 1000
	fmt.Println(mst)
	// struct pointer
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms) // &{42}
	// new function
	var ms2 *myStruct
	ms2 = new(myStruct)
	// (*ms2).foo = 200
	ms2.foo = 200 // you can do it in Go
	fmt.Println(ms2.foo)
	// array and pointer
}

func sayGreeting(greeting, name string) { //synatx suger that writing the same tyoeat the end
	fmt.Println(greeting, name)
}

func sayGreeting2(greeting, name *string) { //synatx suger that writing the same tyoeat the end
	fmt.Println(*greeting, *name)
	*name = "ZZH"
	fmt.Println(*name)
}

func sum(values ...int) {
	fmt.Println(values) //make the variables to a slice
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

func sum2(values ...int) int {
	fmt.Println(values) //make the variables to a slice
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sum3(values ...int) *int {
	fmt.Println(values) //make the variables to a slice
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // return local variable as pointer
}

func sum4(values ...int) (result int) {
	fmt.Println(values) //make the variables to a slice
	for _, v := range values {
		result += v
	}
	return // return variable is decalred on the top
}

func divide(a, b float64) float64 {
	return a / b
}

func divide2(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero!")
	}
	return a / b, nil
}

func funcShow() {
	greeting := "Hello"
	name := "LYC"
	sayGreeting(greeting, name)    // send the copy of variable
	sayGreeting2(&greeting, &name) // passing pointer is more efficient
	fmt.Println(name)
	// 可变参数func
	sum(1, 2, 3, 4, 5)
	// return value
	fmt.Println("The result is ", sum2(1, 2, 3, 4, 5, 6))
	// return a pointer
	fmt.Println("The sum is", *sum3(1, 2, 3, 4, 5, 6, 7))
	// deal with function exception
	fmt.Println(divide(5.0, 3.0))
	fmt.Println(divide(5.0, 0.0))
	// return error
	d, err := divide2(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
	// anonymous function
	/*
		func() {
			fmt.Println("Hello, anonymous function!")
		}
		// error! anonymous function defined, but not be used
	*/
	func() {
		fmt.Println("Hello, anonymous function!")
	}()
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
	// function as type
	var f func() = func() {
		fmt.Println("Hello, f!")
	}
	f()
	f2 := func() {
		fmt.Println("Hello, f2!")
	}
	f2()
	// struct with func, method, OOP?
	g := greeter{
		greeting: "Hello",
		name:     "Method",
	}
	g.greet()
	fmt.Println("The new name is ", g.name)
	g.greet2()
	fmt.Println("The new name is ", g.name)
}

func (g greeter) greet() {
	// copy of struct
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

func (g *greeter) greet2() {
	// pass pointer
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

type greeter struct {
	greeting string
	name     string
}

type Writer interface {
	Write([]byte) (int, error)
}

// implement Write in Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	// 没有出现error接口，go中接口的实现是隐式的
	n, err := fmt.Println(string(data))

	return n, err
}

func interfaceShow() { // 通过interface来实现多态
	// basic
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello GO!"))
	// composing interfaces together
	// type conversion
	// implementing with values vs pointers
	// best practices
}

func main() {
	choice := "channel"

	if choice == "var" {
		varShow()
	} else if choice == "primitives" {
		primeShow()
	} else if choice == "constant" {
		constantShow()
	} else if choice == "Array" {
		arrayShow()
	} else if choice == "Slice" {
		sliceShow()
	} else if choice == "Slice2" {
		sliceShow2()
	} else if choice == "map" {
		mapShow()
	} else if choice == "struct" {
		structShow()
	} else if choice == "flow" {
		flowControl()
	} else if choice == "switch" {
		switchControl()
	} else if choice == "loop" {
		loopShow()
	} else if choice == "defer" {
		// defer1()
		// defer2()
		defer3()
	} else if choice == "panic" {
		// panic1()
		panic2()
	} else if choice == "recover" {
		recover1()
	} else if choice == "pointer" {
		pointerShow()
	} else if choice == "func" {
		funcShow()
	} else if choice == "inter" {
		interfaceShow()
	} else {
		fmt.Println("Hello 世界!")
	}
}
