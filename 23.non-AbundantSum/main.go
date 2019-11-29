package main

import (
	"fmt"
	"math"
)

type primeFac struct {
	prime int
	exp   int
}

func main() {
	fmt.Print("Hello World, enter a number:")
	var x int
	fmt.Scan(&x)
	primeList := setPrimeList(x) //Set the prime List for the number

	perfList := make([]int, 0) //Initialize the abundant List

	for i := 2; i <= x; i++ {
		if abundCheck(i, primeList) { //check every numbers in the loop, from 0 to x
			perfList = append(perfList, i) //only append the abundant list is i is abundant
		}
	}

	remList := make([]int, 0, x) //initialize the list of all combination of 2 abundant numbers

	for i := range perfList { //over the range of abundant numbers
		for j := i; j < len(perfList); j++ { //and all the ones AFTER
			if perfList[i]+perfList[j] < x { //as long as the sum of the two is < x, to stay in scope
				remList = append(remList, perfList[i]+perfList[j]) //appent the combination list
			} else {
				break //or next i
			}
		}
	}

	var remMap = make(map[int]int)

	for i := range remList { //over all the element in a map
		remMap[remList[i]] = 1 //all the elements will be at 1
	}

	sumRem := 0
	for i := range remMap {
		sumRem += i
	}
	sum := 0
	for i := 1; i < x; i++ {
		sum += i
	}

	fmt.Println(sum - sumRem)
	return
}

func abundCheck(a int, primeList []int) bool {
	//first, cover the out of scope posibilities
	if a == 1 {
		return false
	}
	var primeFacList = make([]primeFac, 0, len(primeList)) //create slice of PrimeFac struct
	i, t := a, 0                                           //set i alias of a and t index counter

	for i != 1 { //for each prime divider, as long as divider counter != 1
		if i%primeList[t] == 0 { //as long as prime is a divider
			primeFacList = append(primeFacList, primeFac{primeList[t], 1}) //initialize the prime counter
			i = i / primeList[t]                                           //divide the alias
			for i%primeList[t] == 0 {                                      //as long as this prime i is a divider
				primeFacList[len(primeFacList)-1].exp++ //increment exp count in primeFac struct
				i = i / primeList[t]                    //divide the alias
			}
		} else {
			t++ //else next t when the prime is no longer a divider
		}
	} //until we get to t = 1
	//primeFacList now holds all the prime divisors and their factors

	divLenght := 1
	for i := range primeFacList {
		divLenght = divLenght * (primeFacList[i].exp + 1) //calculate the lenght of entire factor list
	}

	var countList []float64 = make([]float64, divLenght) //create the count list
	for i := range countList {
		countList[i] = 1 //Inialize the count list
	}
	//SET FIRST FACTOR
	var set1 = 1
	for i := range countList {
		countList[i] = math.Pow(float64(primeFacList[0].prime), float64((i % (divLenght / set1) / (divLenght / (primeFacList[0].exp + 1)))))
	}
	set1 *= primeFacList[0].exp + 1

	//COMPLETE THE LIST
	for i := 1; i < len(primeFacList); i++ {

		set1 *= (primeFacList[i].exp + 1)
		lscope := divLenght / set1

		for j := range countList {
			countList[j] *= math.Pow(float64(primeFacList[i].prime), float64((j/lscope)%(primeFacList[i].exp+1)))
		}
	}

	//CONDITION CHECK IF ABUNDANT NUMBER
	var count float64
	for i := 0; i < len(countList)-1; i++ { //over the entire dividor list, except the number itself
		count += countList[i] //you add them
	}
	if int(count) > a { //actual check, is the sum of dividers > number?
		return true //if yes, true
	}
	return false //if no, false
}

func setPrimeList(a int) []int {
	var expLen = (a / 2) + 1
	primeList := make([]int, 2, expLen) //initialize first 2 primes, enough place for half the amount we study
	primeList[0] = 2                    //put first prime in
	primeList[1] = 3                    //put second prime in
	var x int = 5                       //set x to start at next odd after last prime
	for x <= a {
		for i := 0; i < len(primeList); i++ { // testing each candidate x, starting by dividing by 1st prime
			if x%primeList[i] == 0 { //factor found
				x += 2 //next odd x
				break
			} else { //if factor not found
				if primeList[i] >= int(math.Sqrt(float64(x)))+1 { //and we are at square, no factor after, this is a Prime
					primeList = append(primeList, x) //add prime to the list
					x += 2
					break //next odd x
				} else { //next i factor to test
				}
			}
		}
	} //All primes below "a" are now in the slice primeList

	return primeList
}
