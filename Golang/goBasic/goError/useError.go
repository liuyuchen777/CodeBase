package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		return math.Sqrt(f), nil
	}
}

func main() {
	num1 := 4.0
	num2 := -2.2
	if res, err := Sqrt(num1); err == nil {
		fmt.Println("Res is: ", res)
	} else {
		fmt.Println(err)
	}
	if res, err := Sqrt(num2); err == nil {
		fmt.Println("Res is: ", res)
	} else {
		fmt.Println(err)
	}
}