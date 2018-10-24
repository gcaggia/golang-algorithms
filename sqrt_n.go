package main

import (
	"../golang-algorithms/utils"
	"fmt"
	"math"
	"strconv"
)

// We use dichotomy to calculate square root of an integer
func sqrtn(n int, epsilon float64) (int, string) {
	var i, sqrtn_calc, nMin, nMax float64 = 0, 0, 0, float64(n)
	for ; math.Abs(sqrtn_calc * sqrtn_calc - float64(n)) > epsilon; i++ {
		sqrtn_calc = (nMin + nMax) / 2
		if sqrtn_calc * sqrtn_calc > float64(n) {
			nMax = sqrtn_calc
		} else {
			nMin = sqrtn_calc
		}
	}
	return int(i), strconv.FormatFloat(sqrtn_calc, 'f', 10, 64)
}

func main() {
	number := utils.NumberInput("")
	iter, sqrtn := sqrtn(number, 1/1e6)
	fmt.Println("To get a precision of 10^-6, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation sqrt(" +
		strconv.Itoa(number) + ")=" + sqrtn)
}