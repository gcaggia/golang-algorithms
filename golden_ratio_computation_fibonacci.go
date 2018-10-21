package main

import (
	"math"
	"strconv"
	"fmt"
)

func goldenRatioFibonacci(epsilon float64) (int, string) {
	var ratio_calc, i float64
	for fib, prev := float64(3), float64(2);  math.Abs(ratio_calc - math.Phi) > epsilon ;
		fib, prev, i = fib + prev, fib, i + 1 {
		ratio_calc = fib/prev
	}
	return int(i), strconv.FormatFloat(ratio_calc, 'f', 10, 64)
}

func main() {
	iter, phi := goldenRatioFibonacci(1/1e7)
	fmt.Println("To get a precision of 10^-7, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation φ=" + phi)
}