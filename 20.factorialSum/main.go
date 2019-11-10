package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")
	primeMapA := factDecomp(1, 100)

	if primeMapA[2] > primeMapA[5] {
		primeMapA[2] -= primeMapA[5]
		primeMapA[5] = 0
	} else {
		primeMapA[5] -= primeMapA[2]
		primeMapA[2] = 0
	}
	fmt.Println(primeMapA)

	var count int
	for i := range primeMapA {
		count += primeMapA[i]
	}

	var primeSliceA []int
	primeSliceA = make([]int, 0, count)
	for i := range primeMapA {
		for j := 0; j < primeMapA[i]; j++ {
			primeSliceA = append(primeSliceA, i)
		}
	}
	fmt.Println(primeSliceA)

	var lastDigit int = 1
	for i := range primeSliceA {
		lastDigit = (lastDigit * (primeSliceA[i] % 10)) % 10
	}
	fmt.Println(lastDigit)

	var logSum float64
	for i := range primeMapA {
		log := math.Log10(float64(i))
		logSum += float64(primeMapA[i]) * log
	}
	fmt.Println("Total Log10 is", logSum)

	return
}
