package main

import (
	"fmt"
	"strconv"
)

func prime(number int) bool {
	var i int
	for i =2; i < number && number%i != 0; i++ {}
	return i == number
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if prime(number) {
		fmt.Println(strconv.Itoa(number) + " is prime")
	} else {
		fmt.Println(strconv.Itoa(number) + " is not prime")
	}
}
