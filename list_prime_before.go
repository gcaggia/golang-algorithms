package main

import (
	"fmt"
	"strconv"
)


func listPrimesBefore(number int) ( []int) {
	var listPrimes []int
	// var i int
	for n := 2 ; n <= number ; n++ {
		factor := 2
		for ; factor < n && n%factor != 0; factor++ {}
		if factor == n { listPrimes = append(listPrimes, n) }
	}
	return listPrimes
}

func main() {
	var number int
	var listPrimes []int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	listPrimes = listPrimesBefore(number)
	fmt.Print("List of prime numbers before " + strconv.Itoa(number) + " is : ")
	fmt.Println(listPrimes)
}
