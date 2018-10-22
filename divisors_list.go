package main

import (
	"fmt"
	"strconv"
	"../golang-algorithms/utils"
)

func DivisorsList(n int) (divList[] int) {
	for i:= 1; i <=n ; i++ { if n%i == 0 { divList = append(divList, i) } }
	return
}

func main() {
	n:= utils.NumberInput("")
	list := DivisorsList(n)
	fmt.Print("List of divisors of " + strconv.Itoa(n) + ": ", list)
}