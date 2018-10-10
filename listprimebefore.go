package main

import (
	"fmt"
	"strconv"
)

func listPrimesBefore(number int) (int, []int) {
	var listPrimes = [5] int{2, 4, 6, 8, 10}
	// var i int
	// for i = 2; i < number && number%i != 0; i++ {}
	// return i == number
	return number, listPrimes
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if isPrime(number) {
		fmt.Println(strconv.Itoa(number) + " is prime")
	} else {
		fmt.Println(strconv.Itoa(number) + " is not prime")
	}
}
