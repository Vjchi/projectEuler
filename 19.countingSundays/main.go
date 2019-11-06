package main

import "fmt"

func main() {

	yr := 1904

	for i := 0; i < 12; i++ {
		fmt.Println(yr, " - ", i, " - ", (countPastYrs(yr)+countCurYr(yr, i))%7)
	}
	/*var counter int
	for i := 1901; i <= 2001; i++ {
		pastYr := countPastYrs(i)
		for j := 0; j < 12; j++ {
			if (countCurYr(i, j)+pastYr)%7 == 0 {
				counter++
			}
		}
	}
	fmt.Println(counter)*/
	return
}

func countCurYr(yr, month int) int {
	var x int //set the day counter
	x = daysPastMonth[month]
	if month > 1 { //if we are after February
		if yr%400 == 0 { //and the yr is multiple of 400, add 29/2
			x++
			return x //and return the number of days
		}
		if yr%100 == 0 { //but if century, don't add 29/2
			return x //and return the number of days
		}
		if yr%4 == 0 { //but if still multiple of 4, add 29/2
			x++
			return x //and return the number of days
		}
	}
	return x
}

func countPastYrs(yr int) int {
	y := yr*365 + yr/4 - yr/100 + yr/400
	y -= dayOne
	return y
}
