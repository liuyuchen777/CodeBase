package main

import "fmt"

type DivideError struct {
	dividee float64
	divider float64
}

func (de DivideError) Error() string {
	return fmt.Sprintf("Cannot Proceed, the divider is zeror divider: %.2f, divier: 0", de.dividee)
}

func Divide(varDividee, varDivider float64) (float64, error) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		return 0, dData
	} else {
		return varDividee / varDividee, nil
	}
}

func main() {
	dividee1, divider1 := 1., 2.
	dividee2, divider2 := 2.2, 0.
	if res, err := Divide(dividee1, divider1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	if res, err := Divide(dividee2, divider2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

