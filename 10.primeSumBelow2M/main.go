package main

import "fmt"

func main() { //testing the function for all below 1000 first

	primeList := make([]int, 5, 400) //initialize first 5 primes
	primeList[0] = 1
	primeList[1] = 2
	primeList[2] = 3
	primeList[3] = 5
	primeList[4] = 7

	var x int = 9 //set x to start at next odd after last prime

	for x < 1000 {

		for i := 1; i < len(primeList); i++ { // testing each candidate x, starting by dividing by 1st prime

			if x%primeList[i] == 0 { //factor found
				x += 2 //next odd x
				break
			} else { //if factor not found

				if primeList[i] == x/3 { //and we are at third, no factor after, this is a Prime
					primeList = append(primeList, x)
					x += 2 //next odd x
					fmt.Println(primeList)
					break
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
