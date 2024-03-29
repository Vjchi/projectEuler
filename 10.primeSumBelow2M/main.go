package main

import (
	"fmt"
	"math"
)

func main() { //testing the function for all below 1000 first

	primeList := make([]int, 2, 400000) //initialize first 2 primes
	primeList[0] = 2
	primeList[1] = 3

	var x int = 5 //set x to start at next odd after last prime

	for x < 2000000 {

		for i := 1; i < len(primeList); i++ { // testing each candidate x, starting by dividing by 1st prime

			if x%primeList[i] == 0 { //factor found
				x += 2 //next odd x
				break
			} else { //if factor not found

				if primeList[i] >= int(math.Sqrt(float64(x)))+1 { //and we are at third, no factor after, this is a Prime
					primeList = append(primeList, x) //add prime to the list
					x += 2
					break //next odd x
				} else { //next i factor to test
				}
			}
		}
	} //All primes below 2000000 are now in the slice primeList

	var count int                         //create a counter
	for i := 0; i < len(primeList); i++ { //over the whole lenght of the list
		count += primeList[i] //sum all the list
	}

	fmt.Println("The sum of all Primes below 2M is ", count)

	return
}
