package utils

import "fmt"

func NumberInput(txt string) int {
	var number int
	fmt.Print(func() string { if txt == "" { return "Enter a number: " } else {return txt} }())
	fmt.Scan(&number)
	return number
}

func StringInput(txt string) string {
	var input string
	fmt.Print(func() string { if txt == "" { return "Enter a word: " } else {return txt} }())
	fmt.Scanln(&input)
	return input
}