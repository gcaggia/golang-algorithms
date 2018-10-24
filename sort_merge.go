package main

import (
	"fmt"
	utils "../golang-algorithms/utils"
)

func merge(arr[] int, iLeft, iMiddle, iRight int)  {
	// We need a tmp array
	var arrTmp[] int
	// basic variables
	i, j := iLeft, iMiddle + 1

	for k := iLeft; k <= iRight; k++ {
		if i <= iMiddle && j <= iRight {
			if arr[i] < arr[j] {
				arrTmp, i = append(arrTmp, arr[i]), i + 1
			} else {
				arrTmp, j = append(arrTmp, arr[j]), j + 1
			}
		} else {
			if i > iMiddle {
				arrTmp, j = append(arrTmp, arr[j]), j + 1
			} else {
				arrTmp, i = append(arrTmp, arr[i]), i + 1
			}
		}
	}
	for k, l := iLeft, 0; l < len(arrTmp) ; k, l = k + 1, l + 1  {
		arr[k] = arrTmp[l]
	}
}

func mergeSort(arr[] int, iLeft int, iRight int)  {
	if iLeft < iRight {
		var iMiddle int = ( iLeft + iRight ) / 2
		mergeSort(arr, iLeft, iMiddle)
		mergeSort(arr, iMiddle + 1, iRight)
		merge(arr, iLeft, iMiddle, iRight)
	}
}

func main() {
	arr := utils.RandomIntArray(10)
	fmt.Println("Array before: ", arr)
	mergeSort(arr, 0, len(arr) - 1)
	fmt.Println("Array after:  ", arr)
}