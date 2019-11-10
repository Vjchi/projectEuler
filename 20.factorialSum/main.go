package main

import (
	"fmt"
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

	var primeSlA []int
	primeSlA = make([]int, 0, count)
	for i := range primeMapA {
		for j := 0; j < primeMapA[i]; j++ {
			primeSlA = append(primeSlA, i)
		}
	}
	fmt.Println(primeSlA)

	var numCnt []int = make([]int, 5000, 5000)
	var remCnt []int = make([]int, 5000, 5000)
	numCnt[0], remCnt[0] = 1, 0
	for i := range primeSlA {
		for j := range numCnt {
			numCnt[j] = numCnt[j] * primeSlA[i]
			remCnt[j+2] = numCnt[j] / 100
			remCnt[j+1] = numCnt[j] / 10
			numCnt[j] = numCnt[j] % 10
		}
	}

	return
}
