package main

import (
	"fmt"
	"strconv"
	"../golang-algorithms/utils"
)

func PerfectNumberList(n int) (perfectList[] int){
	for i:= 1; i <=n ; i++ {
			sum, max_divisor := 0, int(i/2)
		for divisor:= 1; divisor <= max_divisor ; divisor++ {
			if i%divisor == 0 { sum+=divisor }
			if sum > i { break }
		}
		if sum == i { perfectList = append(perfectList, i)}
	}
	return
}

func main() {
	n:= utils.NumberInput("")
	list := PerfectNumberList(n)
	fmt.Print("List of perfect numbers before " + strconv.Itoa(n) + ": ")
	fmt.Println(list)
}