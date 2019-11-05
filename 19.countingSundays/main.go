package main

import "fmt"

func main() {
	fmt.Println(countPastYrs(1900) - countPastYrs(1899))
	fmt.Println((countPastYrs(2000) + countCurYr(2000, 0)) % 7)
	return
}

func countCurYr(yr, month int) int {
	var x int
	x = daysPastMonth[month]
	if month > 1 {
		if yr%400 == 0 {
			x++
			return x
		}
		if yr%100 == 0 {
			return x
		}
		if yr%4 == 0 {
			x++
			return x
		}
	}
	return x
}

func countPastYrs(yr int) int {
	y := (yr)*365 + (yr)/4 - (yr)/100 + (yr)/400
	y -= dayOne
	return y
}
