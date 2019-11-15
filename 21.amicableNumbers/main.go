package main

import "fmt"

type primeFac struct {
	prime int
	exp   int
}

func main() {
	var primeList []int //Set primeList variable
	fmt.Println("Enter the lenght you'd like to study:")
	var x int                   //set variable to study
	fmt.Scan(&x)                //Get variable from user
	primeList = setPrimeList(x) //create prime List from variable
	sumDiv(x, primeList)
	return
}

func sumDiv(a int, primeList []int) {
	var primeFacList = make([]primeFac, 0, len(primeList)) //create slice
	i, t := a, 0

	for i != 1 { //for each prime divider, as long as alias != 1
		if i%primeList[t] == 0 { //as long as prime is a divider
			primeFacList = append(primeFacList, primeFac{primeList[t], 1}) //initialize the prime counter
			i = i / primeList[t]                                           //divide the alias
			for i%primeList[t] == 0 {
				primeFacList[len(primeFacList)-1].exp++
				i = i / primeList[t]
			}
		} else {
			t++ //else next t
		}
	} //until we get to t = 1
	fmt.Println(primeFacList)
	return
}

func setPrimeList(a int) []int {
	var expLen = (a / 2) + 1
	primeList := make([]int, 2, expLen) //initialize first 2 primes, enough place for half the amount we study
	primeList[0] = 2                    //put first prime in
	primeList[1] = 3                    //put second prime in
	var x int = 5                       //set x to start at next odd after last prime
	for x < a {
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
