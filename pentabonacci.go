package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

func pentafibonnaciSequence(maxIndex int) (seq[][] int) {
	if maxIndex <= 4 { return [][]int{ {0, 0}         } }
	seq = append(seq, [][]int{{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 4}}...)
	for i := 2 ; i <= maxIndex ; i++ {
		seq = append(seq, [][]int{
				{i, seq[i-1][1] + seq[i-2][1] + seq[i-3][1] + seq[i-4][1] + seq[i-5][1] }}...
			)
	}
	return seq
}

func main() {
	t := utils.NumberInput("Enter the max index of fibonacci sequence: ")
	sequence := fibonnaciSequence(t)
	fmt.Println(sequence)
}