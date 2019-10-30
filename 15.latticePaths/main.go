package main

//the solution is the number of possibilities of 20 amongst 40
//which is equal to 40!/(20!20!)
//which my computer is too much of a potato to find on its own
//but I'm right I just checked it on wolfram alpha

func main() {
	var primeList []int = make([]int, 2, 40)
	primeList[0] = 2
	primeList[1] = 3
	x := 5
	for x < 40 {
		for i := 0; i < len(primeList); i++ { // testing each candidate x,
			if x%primeList[i] == 0 { //by dividing by first prime
				x += 2 //next odd x if factor found
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
	} //All primes below 40 are now in the slice primeList
	//now let's put them all in a map
	var primeMap = make(map[int]int) //create map
	for i := range primeList {       //for all the primes identified in primeList
		primeMap[primeList[i]] = 0 //initialise an entry in the map to 0
	} //now all the primes have a counter in the map

	return
}
