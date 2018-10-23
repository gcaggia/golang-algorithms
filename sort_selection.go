package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

func selectionSort(arr[] int)  {
	for i := 0 ; i < len(arr) ; i++  {
		min := i
		for j := i+1; j < len(arr); j++ {
			if arr[j] < min { min = j }
		}
		if min != i { arr[i], arr[min] = arr[min], arr[i] }
	}
}

func main() {
	arr := utils.RandomIntArray(10)
	fmt.Println("Array before: ", arr)
	selectionSort(arr)
	fmt.Println("Array after:  ", arr)
}