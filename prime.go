package main

import (
	"fmt"
	"strconv"
	utils "../golang-algorithms/utils"
)

func isPrime(number int) bool {
	var i int
	for i = 2; i < number && number%i != 0; i++ {}
	return i == number
}

func main() {
	number := utils.NumberInput("")
	if isPrime(number) {
		fmt.Println(strconv.Itoa(number) + " is prime")
	} else {
		fmt.Println(strconv.Itoa(number) + " is not prime")
	}
}
