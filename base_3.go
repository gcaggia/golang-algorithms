package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
	"strconv"
)

func decimalToTernary(number int) (binaryFormat string ) {
	for quotient := number ; quotient != 0 ;  {
		binaryFormat = strconv.Itoa(quotient%3) + binaryFormat
		quotient /=3
	}
	return "(" + binaryFormat + ")"
}

func main()  {
	number := utils.NumberInput("Enter decimal number: ")
	base3 := decimalToTernary(number)
	fmt.Println("Ternary number: " + base3)
}