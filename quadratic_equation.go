package main

import (
	"../golang-algorithms/utils"
	"fmt"
	"strconv"
)

func quadraticEquation(a, b, c int) (solution string)  {
	delta := b * b - 4 * a * c
	if delta < 0 {
		solution = "The discriminant is negative, then there are no real roots"
	} else if delta == 0 {
		solution = "The discriminant is zero, then there is exactly one real root x=-" +
			strconv.Itoa(b) + "/2*" + strconv.Itoa(a)
	} else {
		solution = "The discriminant is positive, then there are two distinct roots: " +
			" x1=-"     + strconv.Itoa(b) + "-sqrt(" + strconv.Itoa(delta) +
			")/2*"      + strconv.Itoa(a) +
			" and x2=-" + strconv.Itoa(b) + "+sqrt(" + strconv.Itoa(delta) +
			")/2*"      + strconv.Itoa(a)
	}
	return
}


func main() {
	fmt.Println("The Standard Form of a Quadratic Equation is ax2 + bx + c = 0")
	fmt.Println("Please choose a, b, c")
	a := utils.NumberInput("a: ")
	b := utils.NumberInput("b: ")
	c := utils.NumberInput("c: ")
	fmt.Println(quadraticEquation(a, b ,c))
}