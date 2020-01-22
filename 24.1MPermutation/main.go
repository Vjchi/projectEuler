package main

import (
	"fmt"
)

func main() {

	//First create the permutation list and result slice
	list := make([]int, 0, 10)
	result := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		list = append(list, i)
	}
	fmt.Println(list)

	//Then setup the algorithgm
	//Initialize the first conditions (999999 because index strats at 0)

	objective := 999999

	for i := len(list); i > 0; i-- {

		//First calculate the stage size
		var size int = 1
		for j := i - 1; j > 1; j-- {
			size *= j
		}
		fmt.Println("Stage ", i, ", the stage size is ", size)

		//Calculate the instances of sizes over the objective
		rest := objective / size

		//Amend result Slice, original list and remmaining objective
		result = append(result, list[rest])
		list = remIndex(list, rest+1)
		objective = objective - rest*size

		fmt.Println("result so far: ", result, "-- list is", list, "-- obj rem is", objective)
	}
	//Print result
	fmt.Println("The finish slice is ", result)

	return
}

func remIndex(slice []int, index int) []int {
	retSlice := make([]int, 0, len(slice))

	if index == len(slice) {
		retSlice = slice[:len(slice)-1]
	} else {
		retSlice = append(slice[0:index-1], slice[index:]...)
	}
	return retSlice
}
