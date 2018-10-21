package utils

import "fmt"

func NumberInput(txt string) int {
	var number int
	fmt.Print(func() string { if txt == "" { return "Enter a number: " } else {return txt} }())
	fmt.Scan(&number)
	return number
}