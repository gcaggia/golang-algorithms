package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

func insertionSort(arr[] int)  {
	for i := 1; i < len(arr) ; i++  {
		for j:= i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
 		}
	}
}

func main() {
	arr := utils.RandomIntArray(10)
	fmt.Println("Array before: ", arr)
	insertionSort(arr)
	fmt.Println("Array after:  ", arr)
}