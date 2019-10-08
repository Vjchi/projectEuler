package main

import "fmt"

func main() {

	var count, x int = 6, 15 //6th prime is 13, so start counting at 5 and testing at 13

	for i := 2; i < (x/2)+1; i++ { // testing each candidate x, starting by dividing by 2
		if x%i == 0 { //if factor found
			x += 2 //next x
			i = 2  // re initialize i

		} else { //if factor not found
			if i == x/2 { //and we are at half, no factor after
				count++ //we found another prime, count it

				if count == 10001 { //when 10001 prime is found
					fmt.Println("The ", count, "th prime is ", x) //present result
					return                                        //finish function

				} else { //if not Prime No.10001
					x += 2 //next x
					i = 2  // re initialize i
				}

			} else { //next i factor to test
			}
		}
	}
	fmt.Println("The ", count, "th prime is ", x)
	return
}
