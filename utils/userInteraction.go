package utils

import "fmt"

func numberInput() int {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	return number
}