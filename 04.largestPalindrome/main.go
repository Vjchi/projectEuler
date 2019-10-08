package main

import (
	"fmt"
	"math"
)

func main() {

	//var x int //this is to test palindrome
	//fmt.Scan(&x) //scan a value
	//fmt.Println(palindrome(x)) //check if palindrome works

	//var count int = 999

	//	for i := 999; i > 99; i-- {
	//	check := count * i
	//	if palindrome(count * i) {

	//	} else {

	//	}

	//	}
	reversePalindrome()

	return
}

func palindrome(numb int) bool {

	var a, b, c, d, e, f int //put the digit variables in place
	var a1, b1, c1, d1 int   //put the substract variables in place

	a = numb / 100000
	a1 = numb - a*100000

	b = a1 / 10000
	b1 = a1 - b*10000

	c = b1 / 1000
	c1 = b1 - c*1000

	d = c1 / 100
	d1 = c1 - d*100

	e = d1 / 10
	f = d1 - e*10

	fmt.Println(a, b, c, d, e, f)

	if a == f && b == e && c == d {
		return true
	} else {
		return false
	}
}

func reversePalindrome() {
	var a, b, c int = 9, 9, 9

	for a >= 0 {

		for b >= 0 {

			for c > 0 {
				x := a*100001 + b*10010 + c*1100

				var y, z int = checkFactor(x)

				fmt.Println(y, z)

				if y > 99 {
					fmt.Println("The two factors are ", y, " and ", z, " and their product is ", x)
					return
				} else {

				}
				c--
			}
			b--
			c = 9
		}
		a--
		b = 9
	}
	return
}

func checkFactor(x int) (a, b int) {

	var y, z int

	conv := float64(x)

	u := math.Sqrt(conv) + 1
	y = int(u)

	for y > 900 {

		if x%y == 0 {
			z = x / y
			if z < 1000 && z > 99 {
				return y, z
			} else {
			}
		} else {
		}
		y--
	}
	return 0, 0
}
