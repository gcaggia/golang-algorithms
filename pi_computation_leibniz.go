package main

import (
	"math"
	"strconv"
	"fmt"
)
// Link: https://en.wikipedia.org/wiki/Leibniz_formula_for_%CF%80
func PiComputationLeibniz(epsilon float64) (int, string) {
	var pi_calc, i float64
	for ;  math.Abs(pi_calc - math.Pi) > epsilon ; i++ {
		pi_calc += 4 * math.Pow(-1, float64(i)) / (2 * i + 1)
	}
	return int(i), strconv.FormatFloat(pi_calc, 'f', 10, 64)
}

func main() {
	iter, pi := PiComputationLeibniz(1/1e7)
	fmt.Println("To get a precision of 10^-7, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation π=" + pi)
}