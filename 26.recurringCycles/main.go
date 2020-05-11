package main

import "fmt"

func main() {

	var result int
	var longestChain []int
	longestChain = make([]int, 0, 1000)

	for i := 2; i < 100; i++ {
		a := recurringCycle(i)
		if len(a) > len(longestChain) {
			longestChain = a
			result = i
		}
	}
	fmt.Println("The longest recurring cycle is", longestChain, " and happen for d =", result, "and lenght is ", len(longestChain))
	return
}

func recurringCycle(a int) []int {
	var cycle, addedCycle, restCycle []int
	cycle = make([]int, 0, 105)
	restCycle = make([]int, 0, 105)

	var rest int = 0
	var limit int

	rest, addedCycle = divide(1, a)
	cycle = append(cycle, addedCycle...)
	restCycle = append(restCycle, rest)

	for rest != 0 && rest != 1 && limit < 100 {
		rest, addedCycle = divide(rest, a)
		cycle = append(cycle, addedCycle...)

		for i := range restCycle {
			if rest == restCycle[i] {
				limit = 1000 //ca suffit pas, je tombe sur une récurrence irréelle
			}
		}

		limit++
	}
	return cycle
}

func divide(num int, div int) (int, []int) {

	var rest int
	var result []int

	result = make([]int, 0, 4)

	num *= 10
	if div > num {
		for div > num {
			result = append(result, 0)
			num *= 10
		}
	}

	rest = num % div
	result = append(result, (num-rest)/div)

	return rest, result
}
