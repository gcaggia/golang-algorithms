package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

func fibonnaciSequence(maxIndex int) (seq[][] int) {
	if maxIndex <= 0 { return [][]int{ {0, 1}         } } else
	if maxIndex == 1 { return [][]int{ {0, 1}, {1, 1} } }
	seq = append(seq, [][]int{{0, 1}, {1, 1}}...)
	for i := 2 ; i <= maxIndex ; i++ {
		seq = append(seq, [][]int{{i, seq[i-1][1] + seq[i-2][1]}}...)
	}
	return seq
}

func main() {
	t := utils.NumberInput("Enter the max index of fibonacci sequence: ")
	sequence := fibonnaciSequence(t)
	fmt.Println(sequence)
}