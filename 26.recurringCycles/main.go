package main

import "fmt"

func main() {

	var result int
	var longestChain []int
	longestChain = make([]int, 0, 100)

	for i := 1; i < 10; i++ {
		a := recurringCycle(i)
		if len(a) > len(longestChain) {
			longestChain = a
			result = i
		}
	}
	fmt.Println("The longest recurring cycle is", longestChain, " and happen for d =", result)
	return
}

func recurringCycle(a int) []int {
	var cycle, addedCycle []int
	cycle = make([]int, 0, 100)

	var rest int
	var limit int

	rest, addedCycle = divide(1, a)
	cycle = append(cycle, addedCycle...)

	for rest != 0 && rest != 1 && limit < 100 {

		rest, addedCycle = divide(rest, a)
		cycle = append(cycle, addedCycle...)
		limit++
		fmt.Println("Rest is ", rest)
	}
	fmt.Println("Test 1 -", a, "-", cycle)
	return cycle
}

func divide(num int, div int) (int, []int) {

	var rest int
	var result []int

	result = make([]int, 0, 4)

	for div > num {
		result = append(result, 0)
		num *= 10
	}

	rest = num % div
	result = append(result, (num-rest)/div)

	return rest, result
}
