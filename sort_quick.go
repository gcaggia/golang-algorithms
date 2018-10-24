package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

func pivot(arr[] int, iLeft int, iRight int) (iPivot int) {
	iPivot = iLeft
	for i, valInitPivot :=iLeft + 1, arr[iLeft] ; i <= iRight; i++  {
		if arr[i] < valInitPivot {
			iPivot++
			arr[i], arr[iPivot] = arr[iPivot], arr[i]
		}
	}
	arr[iPivot], arr[iLeft] = arr[iLeft], arr[iPivot]
	return
}

func quickSort(arr[] int, iLeft int, iRight int)  {
	if iLeft < iRight {
		iPivot := pivot(arr, iLeft, iRight)
		quickSort(arr, iLeft, iPivot)
		quickSort(arr, iPivot + 1, iRight)
	}
}

func main() {
	arr := utils.RandomIntArray(10)
	fmt.Println("Array before: ", arr)
	quickSort(arr, 0, len(arr) - 1)
	fmt.Println("Array after:  ", arr)
}