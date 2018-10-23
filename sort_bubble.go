package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

func bubbleSort(arr[] int)  {
	for  flag := true ; flag == true ;  {
		flag = false
		for i := 0; i < len(arr) - 1; i++ {
			if arr[i] > arr[i+1] { arr[i], arr[i+1], flag = arr[i+1], arr[i], true }
		}
	}
}

func main() {
	arr := utils.RandomIntArray(10)
	fmt.Println("Array before: ", arr)
	bubbleSort(arr)
	fmt.Println("Array after:  ", arr)
}