package main

import "fmt"

func main() {

	var primeList = make([]int, 3, 4000)
	primeList[0] = 2 //first two primes in the list
	primeList[1] = 3 //otherwise it fucks up the program
	primeList[2] = 5 //like, really badly

	var x int = 7 //following check is the next odd number after last prime

	for len(primeList) < 10001 {

		var i int = 0 //re intialize i

		for i < len(primeList) {
			if x%primeList[i] == 0 { //if factor found
				x += 2 //next x
				break  //break the loop

			} else { //if factor not found
				if primeList[i] > x/3 { //and we are at third, no factor after
					primeList = append(primeList, x) //we found another prime, append it
					if len(primeList) == 10001 {
						break //10001 we'll stop here
					}
					x += 2 //otherwise next odd number will be checked
					break  //re initialize i
				} else {
					i++ //if we not a third, next i
				}
			}
		}
	}
	fmt.Println("The early ", len(primeList), "th prime is ", x) //x-2 because x is incremented before end
	return
}
