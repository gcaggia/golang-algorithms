package main

import (
	"fmt"
	"../golang-algorithms/utils"
)

func PascalTriangle(deep int) (triangle [][]int) {
	for line := 0; line <= deep; line++ {
		triangleLine := []int{}
		for elem := 0; elem <= line ; elem++ {
			if elem == 0 || line == elem {
				triangleLine = append(triangleLine, 1)
			} else {
				triangleLine = append(triangleLine, triangle[line-1][elem-1] + triangle[line-1][elem])
			}
		}
		triangle = append(triangle, triangleLine)
	}
	return triangle
}

func main() {
	n := utils.NumberInput("How many rows for this pascal triangle ? ")
	triangle := PascalTriangle(n)
	fmt.Println("Here is your pascal triangle: ")
	for _, row := range triangle { fmt.Println(row) }
}