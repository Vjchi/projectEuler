package main

import (
	"fmt"
)

func main() {
	fmt.Println("Type in the factorial to which you want to calculate the digit sum:")
	fmt.Println("(limit is 1000)")
	var x int
	fmt.Scan(&x)

	if x > 1000 {
		fmt.Println("Required number is too high")
		fmt.Println("System overload. Emergency shutdown")
		return
	}

	primeMapA := primeDecomp(x)

	if primeMapA[2] > primeMapA[5] {
		primeMapA[2] -= primeMapA[5]
		primeMapA[5] = 0
	} else {
		primeMapA[5] -= primeMapA[2]
		primeMapA[2] = 0
	}

	var count int
	for range primeMapA {
		count++
	}

	var primeSlA []int
	primeSlA = make([]int, 0, count)
	for i := range primeMapA {
		for j := 0; j < primeMapA[i]; j++ {
			primeSlA = append(primeSlA, i)
		}
	}
	fmt.Println(primeSlA)
	var numCnt []int = make([]int, 200, 200) //set our digit list
	numCnt[0] = 1                            //initialise it

	for i := range primeSlA {
		var remCnt []int = make([]int, 203, 203) //set our remainder list
		for j := range numCnt {
			numCnt[j] = numCnt[j] * primeSlA[i] //multiply each digit by the prime
			remCnt[j+2] += numCnt[j] / 100      //Store remainder >100
			numCnt[j] = numCnt[j] % 100         //remove remainder part >100
			remCnt[j+1] += numCnt[j] / 10       //Store remainder >10
			numCnt[j] = numCnt[j] % 10          //Only keep unit digit
		} //next j

		for j := range numCnt { //now to associate remainder to counter
			numCnt[j] += remCnt[j] //add the remainder to our digit counter
		} //next j

		for j := 0; j < len(numCnt)-2; j++ { //now to rearrange counter
			numCnt[j+2] += numCnt[j] / 100 //add count >100
			numCnt[j] = numCnt[j] % 100    //remove count part >100
			numCnt[j+1] += numCnt[j] / 10  //add count >10
			numCnt[j] = numCnt[j] % 10     //only keep singl digit
		} //next j
	} //by the end of it, we should have all the digits making 100!
	var total int
	for i := range numCnt {
		total += numCnt[i]
	}
	fmt.Println(numCnt[0:200]) //print the complete number, in the wrong order
	fmt.Println(total)

	return
}
