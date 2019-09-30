package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)
	fmt.Println(palindrome(x))
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
