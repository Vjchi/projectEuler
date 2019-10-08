package main

import "fmt"

func main() {

	var count, count2 int

	for i := 1; i <= 100; i++ {
		count += i * i
		count2 += i
		fmt.Println("CHECK", i)
	}

	count2 = count2 * count2
	diff := count2 - count

	fmt.Println("Sum of the squares is ", count, " and the square of the sum is ", count2)
	fmt.Println("And the difference is ", diff)

	return
}
