package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
	"strconv"
)

var hexa = map[int]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5",
	6: "6", 7: "7", 8: "8", 9: "9",
	10:"A", 11:"B", 12:"C", 13:"D", 14:"E", 15:"F",
}

func decimalToBaseN(number, base int) (baseN[] string ) {
	for quotient := number ; quotient != 0 ;  {
		if base == 16 {
			baseN = append([]string{hexa[quotient%base]}, baseN... )
		} else {
			baseN = append([]string{strconv.Itoa(quotient%base)}, baseN... )
		}
		quotient /=base
	}
	return
}

func main()  {
	number := utils.NumberInput("Enter decimal number: ")
	base := utils.NumberInput("Choose base: ")
	baseN := decimalToBaseN(number, base)
	fmt.Print("Base " + strconv.Itoa(base) + " number: ")
	fmt.Println(baseN)
}