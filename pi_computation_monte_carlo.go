package main

import (
	"math"
	"strconv"
	"fmt"
	"math/rand"
	"../golang-algorithms/utils"
)

// 4 * lim n -> +inf ( sigma Pn ∈ Quadrant / P_total ) = π
func PiComputationMonteCarlo(epsilon float64) (int, string) {
	var Pn, i float64 = 0, 1
	for ;  math.Abs(4 * Pn/i - math.Pi) > epsilon ; i++ {
		x, y := rand.Float64(), rand.Float64()
		Pn += utils.Bool2Float( (math.Pow(x, 2) + math.Pow(y, 2) ) < 1 )
	}
	return int(i), strconv.FormatFloat(4 * Pn/i , 'f', 10, 64)
}

func main() {
	iter, pi := PiComputationMonteCarlo(1/1e7)
	fmt.Println("To get a precision of 10^-7, we needed " +
		strconv.Itoa(iter) + " iterations and with the computation π=" + pi)
}