package main

import (
	"fmt"
)

func main() {

	var a, b, c float32

	for b = 1; b < 500; b++ {
		a = (500 - b) / (1 - 0.001*b)
		check := float32(int(a)) //create check, natural part of a
		if a-check != 0 {        //look for decimal points
			//if decimal point, not int, next b
		} else {
			c = 1000 - b - a
			fmt.Println("The triplet was found, it's ", a, ", ", b, "and ", c)
			fmt.Println("and their product is ", a*b*c)
			return
		}
	}
	return
}
