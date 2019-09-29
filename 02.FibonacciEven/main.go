package main

import "fmt"

func main() {

	var i, iminus, total int = 1, 0, 0

	for i < 4000000 {
		if i%2 == 0 {
			total += i
		} else {
		}
		i += iminus
		iminus = i
	}
	fmt.Println(total)
}
