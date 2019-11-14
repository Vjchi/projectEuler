package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var primeList []int
	primeList = setPrimeList()
	sumDiv(4856, primeList)
	return
}

func sumDiv(a int, primeList []int) {
	var primeMap = make(map[int]int) //create map
	for i := range primeList {       //for all the primes identified in primeList
		primeMap[primeList[i]] = 0 //initialise prime entries in the map
	} //now all the primes have a counter in the map set to zero
	//now to decompose a in primal terms
	i, t := a, 0
	for i != 1 { //for each prime divider, as long as alias != 1
		if i%primeList[t] == 0 { //if prime is a divider
			primeMap[primeList[t]]++ //increment counter in the map
			i = i / primeList[t]     //divide the alias
		} else {
			t++ //else next t
		} //incremented all the t for the alias
	} //until we get to t = 1
	for i := range primeMap {
		if primeMap[i] == 0 {
			delete(primeMap, i)
		}
	}
	fmt.Println(primeMap)
	return
}

func setPrimeList() []int {
	primeList := make([]int, 2, 2500) //initialize first 2 primes, enough place for 2500 prime
	primeList[0] = 2
	primeList[1] = 3
	var x int = 5 //set x to start at next odd after last prime
	for x < 5000 {
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
	} //All primes below 5000 are now in the slice primeList
	return primeList
}
