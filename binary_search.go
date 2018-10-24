package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
	"math/rand"
	"sort"
	"strconv"
)

func binarySearch(arr[] int, element int) (iteration, position int)  {
	iMin, iMax := 0, len(arr) - 1
	for ; arr[position] != element &&  iMin <= iMax ; iteration++ {
		position = (iMin + iMax)/2
		if arr[position] < element {
			iMin = position
		} else {
			iMax = position
		}
	}
	if iMin > iMax { return iteration,-1 } else { return }
}

func main() {
	size := 10
	arr := utils.RandomIntArray(size)
	sort.Ints(arr)
	valueToFind := arr[rand.Intn(size)]
	fmt.Println("Array before: ", arr)
	fmt.Println("value to find: ", valueToFind)
	iteration, position := binarySearch(arr, valueToFind)
	fmt.Println("Position: " + strconv.Itoa(position) +
		" / nb of iterations: " + strconv.Itoa(iteration))
}