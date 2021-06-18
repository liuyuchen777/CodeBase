package main

import (
	"fmt"
	"sort"
)

func myFunc(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func sortMap(mp map[int]string) []int {
	keys := make([]int, 0, len(mp))

	for i, _ := range mp {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	return keys
}

func main() {
	//myFunc(1, 2, 3)
	//myFunc("liuyuchen", "zhujunyi", 4, 5)
	//
	//MyPrintf(1, "liuyuchen", int64(98))

	//arr := []int{1, 2, 3, 4, 5}
	//for _, v := range arr {
	//	defer fmt.Println(v)
	//}

	// sort a map
	//myMap := make(map[int]string)
	//myMap[4] = "liuyuchen"
	//myMap[7] = "zhujunyi"
	//myMap[8] = "zhoutianyi"
	//
	//fmt.Println("Before Sort: ")
	//for key, val := range myMap {
	//	fmt.Println(key, val)
	//}
	//
	//keys := sortMap(myMap)
	//
	//fmt.Println("After Sort: ")
	//for _, i := range keys {
	//	fmt.Println(i, myMap[i])
	//}

	// sort a string
	//str := []byte("aksdjhcviuhsrfnlcnfe")
	//sort.Slice(str, func (i, j int) bool {
	//	return str[i] < str[j]
	//})
	//fmt.Println(string(str))

}
