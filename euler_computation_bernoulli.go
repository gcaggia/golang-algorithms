package main

import (
	"math"
	"strconv"
	"fmt"
)

// https://en.wikipedia.org/wiki/E_(mathematical_constant)
func EulerBernoulli(epsilon float64) (int, string) {
	var e_calc, n float64 = 0, 1
	for ;  math.Abs(e_calc - math.E) > epsilon ; n++ {
		e_calc = math.Pow(1 + 1/n, n)
	}
	return int(n), strconv.FormatFloat(e_calc, 'f', 10, 64)
}

func main() {
	iter, e := EulerBernoulli(1/1e7)
	fmt.Println("To get a precision of 10^-7, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation e=" + e)
}