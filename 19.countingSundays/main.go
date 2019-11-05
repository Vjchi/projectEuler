package main

import "fmt"

func main() {

	fmt.Println("Hello World!")
	fmt.Println(countPastYrs(1952) - countPastYrs(1951))

	return
}

func countCurYr(yr, month, day int) int {
	var x int

	return x
}

func countPastYrs(yr int) int {
	y := (yr)*365 + yr/4 - yr/100 + yr/400
	y -= dayOne
	return y
}
