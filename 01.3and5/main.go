package main

import "fmt"

func main() {

	var total int

	for i := 0; i < 1000; i++ {
		if i%5 == 0 {
			total += i
		} else if i%3 == 0 {
			total += i
		} else {
		}
	}
	fmt.Println(total)
}
