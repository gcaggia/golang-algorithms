package main

import (
	"math"
	"strconv"
	"fmt"
)

// ln(2) = sum n=1 to inf (-1)^n+1 / n
func ln2(epsilon float64) (int, string) {
	var ln2, ln2_calc, n float64 = math.Log(2), 0, 1
	for ;  math.Abs(ln2_calc - ln2) > epsilon ; n++ {
		ln2_calc += math.Pow(-1, n+1)/n
	}
	return int(n), strconv.FormatFloat(ln2_calc, 'f', 10, 64)
}

func main() {
	iter, ln2 := ln2(1/1e7)
	fmt.Println("To get a precision of 10^-7, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation ln(2)=" + ln2)
}