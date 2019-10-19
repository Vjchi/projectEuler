package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	primeList := generatePrimeList()

	y := TriangularNo(x)
	fmt.Println("The ", x, "th triangular number is ", y)
	fmt.Println("And it has ", factorCount(y, primeList), " factors")

	return
}

func factorCount(x int, list []int) int {
	var i int = 0
	count := make([]int, 1000, 1000)

	for i < len(list) {

		if x%list[i] == 0 { //try to divide it by prime
			x = x / list[i] //if dividable, remove the factorization from the x,
			count[i]++      //increment prime exponent, and try again
		} else { // if not dividable, try next prime
			i++
		}

	}
	var div int = 1
	for j := 0; j < len(count); j++ {
		div = div * (count[j] + 1)
	}
	return div
}

func TriangularNo(x int) int {
	var count int
	if x%2 == 0 {
		count = x*((x+1)/2) + x/2 //That's the real expression of a triangular No.
		return count
	} else {
		count = x * ((x + 1) / 2) //That's the real expression of a triangular No.
		return count
	}
}

func generatePrimeList() []int {

	primeList := make([]int, 2, 1000) //initialize first 5 primes
	primeList[0] = 2
	primeList[1] = 3

	var x int = 5 //set x to start at next odd after last prime

	for x < 1000 {

		for i := 1; i < len(primeList); i++ { // testing each candidate x, starting by dividing by 1st prime

			if x%primeList[i] == 0 { //factor found
				x += 2 //next odd x
				break
			} else { //if factor not found

				if primeList[i] >= x/3 { //and we are at third, no factor after, this is a Prime
					primeList = append(primeList, x) //add prime to the list
					x += 2
					break //next odd x
				} else { //next i factor to test
				}
			}
		}
	} //All primes below 2000000 are now in the slice primeList
	return primeList
}
