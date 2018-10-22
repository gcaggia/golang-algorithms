package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

// https://www.mathsisfun.com/prime-factorization.html
func primeFactorDecomposition(number int) (listFactor[] int) {
	for p := number ; p != 1 ;  {
		for i := 2; i <= p ; i++ {
			if p%i == 0 {
				listFactor = append(listFactor, i)
				p /= i
				break
			}
		}
	}
	return
}

func main()  {
	number := utils.NumberInput("")
	listFactor := primeFactorDecomposition(number)
	fmt.Println(listFactor)
}