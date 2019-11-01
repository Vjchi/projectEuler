package main

import (
	"fmt"
	"math"
)

//the solution is the number of possibilities of 20 amongst 40
//which is equal to 40!/(20!20!)
//which my computer is too much of a potato to find on its own
//but I'm right I just checked it on wolfram alpha

func main() {

	primeMapA := factDecomp(2, 20)  //set the map for A
	primeMapB := factDecomp(21, 40) //set the map for B

	//now let's divide B/A in primeMapC
	primeMapC := make(map[int]int)
	var counter float64 = 1

	for i := range primeMapB { //using the range of the bigger map
		primeMapC[i] = primeMapB[i] - primeMapA[i] //C is the substract of A to B
		counter = counter * math.Pow(float64(i), float64(primeMapC[i]))
	} //and the actual number is just the prime to the power of its counter
	fmt.Println(primeMapC)
	fmt.Println(counter)
	return
}

func factDecomp(a int, b int) map[int]int {
	var primeList []int = make([]int, 2, b)
	var temp int = 5
	primeList[0] = 2
	primeList[1] = 3

	for temp <= b {
		for i := 0; i < 40; i++ { // testing each candidate temp,
			if temp%primeList[i] == 0 { //by dividing by first prime
				temp += 2 //next odd x if factor found
				break
			} else { //if factor not found
				if primeList[i] >= temp/3 { //we are at third, no factor after,
					primeList = append(primeList, temp) //add prime to the list
					temp += 2
					break //next odd temp
				} else { //next i factor to test
				}
			}
		}
	} //All primes below b are now in the slice primeList

	//now let's put them all in a map
	var primeMap = make(map[int]int) //create map
	for i := range primeList {       //for all the primes identified in primeList
		primeMap[primeList[i]] = 0 //initialise any entry in the map to 0
	} //now all the primes have a counter in the map

	//now to decompose (b-a)! in primal terms
	for i := a; i <= b; i++ { //for each factor
		j := i //set an alias
		var t int
		for j != 1 { //for each prime divider, as long as alias != 1
			if j%primeList[t] == 0 { //if prime is a divider
				primeMap[primeList[t]]++ //increment counter in the map
				j = j / primeList[t]     //divide the alias
				t = 0                    //reevaluate if t > 1, reinitialise factors
			} else {
				t++ //else next t
			} //incremented all the t for the alias
		} //until we get to t = 1
	} //increment i, until we get to the last term of factorial
	return primeMap
}
