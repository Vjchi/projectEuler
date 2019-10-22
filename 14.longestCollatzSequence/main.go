package main

import (
	"fmt"
)

func main() {
	var count, maxCount, maxX int //initialize our counters and Maximums

	for x := 1; x < 1000000; x++ { //set up a loop to test all x from 1 to 1 000 000
		a := x    //set an alias for x, the one that will change
		count = 1 //set the count sequence to 1 (for the starting number)

		for {
			if a == 1 { //if a is 1, end of the sequence, break the loop
				break

			} else if a%2 == 0 { //if a is even
				a = a / 2 //apply the formula
				count++   //increment the count

			} else { //if a is odd, and different to one
				a = 3*a + 1 //apply formula
				count++     //increment the count
			}

		}
		if count > maxCount { //test if we have longer sequence
			maxCount = count //update lenght
			maxX = x         //update starting number
		}
	}

	fmt.Println("The longest Collatz lenght is ", maxCount, " and it originated from ", maxX)
	return
}
