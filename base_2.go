package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
	"strconv"
)

func decimalToBinary(number int) (binaryFormat string ) {
	for quotient := number ; quotient != 0 ;  {
		binaryFormat = strconv.Itoa(quotient%2) + binaryFormat
		quotient /=2
	}
	return "(" + binaryFormat + ")"
}

func main()  {
	number := utils.NumberInput("Enter decimal number: ")
	base2 := decimalToBinary(number)
	fmt.Println("Binary number: " + base2)
}