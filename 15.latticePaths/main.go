package main

import (
	"fmt"
	"math"
)

func main() {
	var count, i float64

	for i = 0; i <= 20; i++ {
		count += math.Pow(2, i)
	}
	fmt.Println(count)
	return
}
