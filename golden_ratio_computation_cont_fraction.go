package main

import (
	"math"
	"strconv"
	"fmt"
)

// https://en.wikipedia.org/wiki/Golden_ratio
func goldenRatioContinueFraction(epsilon float64) (int, string) {
	var ratio_calc, i float64 = 1, 1
	for ;  math.Abs(ratio_calc - math.Phi) > epsilon ; i++ {
		ratio_calc = 1 + 1/ratio_calc
	}
	return int(i), strconv.FormatFloat(ratio_calc, 'f', 10, 64)
}

func main() {
	iter, phi := goldenRatioContinueFraction(1/1e9)
	fmt.Println("To get a precision of 10^-9, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation φ=" + phi)
}