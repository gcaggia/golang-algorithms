package utils

import (
	"math/rand"
	"time"
)

func Bool2Float(b bool) float64 {
	if b { return 1 }
	return 0
}

func RandomIntArray(size int) (arr[] int) {
	for i := 0; i< size;i++  {
		rand.Seed(time.Now().UTC().UnixNano())
		n:= rand.Intn(1000) - rand.Intn(1000)
		arr = append(arr, n)
	}
	return
}