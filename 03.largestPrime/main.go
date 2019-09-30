package main

import "fmt"

func main() {
	//biggest prime divider of 600851475143
	var numb int = 600851475143
	//var test int
	//fmt.Scan(&test)
	//fmt.Println(prime(test)) // this is just for testing the prime func
	fmt.Println(numb / 6)
	for i := 1; i <= numb/2; i++ { //limit the scope of search to numb / 2

		if numb%i == 0 { // chech if i is a factor

			t := numb / i         // define opposite factor
			if prime(t) == true { //test if opposite factor is prime number

				fmt.Println("The biggest prime factor of 600851475143 is ", t) //success
				return                                                         //end function

			} else { // if not prime, ignore
			}
		} else { // if not a factor, ignore,
		}
	} //next i
	fmt.Println("Failed") //If I can't find one this is a failure
	return
}

func prime(x int) bool {
	var counter int = x / 2 // define a limit to stop counting after
	for i := 2; i < counter; i++ {
		if x%i == 0 {
			return false
		} else {
		}
	}
	return true
}
