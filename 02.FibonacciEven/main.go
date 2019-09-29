package main

import "fmt"

func main() {

	var inext, i, iminus, total int = 1, 1, 0, 0

	for i < 4000000 {
		if i%2 == 0 {
			total += i
		} else {
		}
		inext = i + iminus
		iminus = i
		i = inext
		fmt.Println("i is ", inext, ", iminus is ", iminus)
	}
	fmt.Println(total)
}
