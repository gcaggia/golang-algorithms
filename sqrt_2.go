package main

import (
	"math"
	"strconv"
	"fmt"
)

// https://en.wikipedia.org/wiki/List_of_mathematical_constants
// sqrt(2) = prod[n=1 to ∞] {1+(-1)^(n+1) /(2n-1)}
func sqrt2(epsilon float64) (int, string) {
	var sqrt2_calc, n float64 = 1, 1
	for ;  math.Abs(sqrt2_calc - math.Sqrt2) > epsilon ; n++ {
		sqrt2_calc *= 1 +  math.Pow(-1, n+1)/ (2*n-1)
		fmt.Println(sqrt2_calc)
	}
	return int(n), strconv.FormatFloat(sqrt2_calc, 'f', 10, 64)
}

func main() {
	iter, sqrt2 := sqrt2(1/1e7)
	fmt.Println("To get a precision of 10^-7, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation sqrt(2)=" + sqrt2)
}