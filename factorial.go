package main

import (
	"fmt"
	"strconv"
	"../golang-algorithms/utils"
)

func factorial(n int) (fact int) {
	if n <= 1 { return 1 } else { return n * factorial(n-1) }
}

func main() {
	number := utils.NumberInput("")
	fmt.Println(strconv.Itoa(number) + "!=" + strconv.Itoa(factorial(number)))
}