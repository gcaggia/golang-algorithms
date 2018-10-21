package main

import (
	"math"
	"strconv"
	"fmt"
)

// https://en.wikipedia.org/wiki/E_(mathematical_constant)
func EulerInfiniteSeries(epsilon float64) (int, string) {
	var e_calc, i float64 = 0, 1
	for fact := 1;  math.Abs(e_calc - math.E) > epsilon ; i, fact = i + 1, fact * int(i)  {
		e_calc += (1/float64(fact))
	}
	return int(i), strconv.FormatFloat(e_calc, 'f', 10, 64)
}

func main() {
	iter, e := EulerInfiniteSeries(1/1e7)
	fmt.Println("To get a precision of 10^-7, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation e=" + e)
}