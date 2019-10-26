package main

import "fmt"

func main() {

	var letterTable map[int]int = initTable()

	var a, total int

	for i := 1; i < 20; i++ { //a counts 1-19
		a += letterTable[i] //capture as a
	} //a is sum of all letters for 1-19
	for i := 20; i < 100; i++ { //now for 20-99
		a += letterTable[(i/10)*10] + letterTable[i%10]
	}

	total = a //a & total are now all letters from 1 to 99

	for i := 1; i < 10; i++ {
		total += 100*(letterTable[i]+letterTable[100]) + 99*3 + a
	} //adding hundreath plus all the letters between 1-99
	//99*3 is for the "and" mentionned 99 times

	total += 11           //finally add the 1000
	fmt.Println(a, total) //Print the total

	/*fmt.Println("What number to test?")
	var x int
	fmt.Scan(&x)
	fmt.Println(x, " has ", countLetter(x), " letters")*/

	return
}

func countLetter(x int) int {
	var letterTable map[int]int = initTable()
	if x <= 20 {
		return letterTable[x]
	}

	if x%100 >= 20 {
		return letterTable[x/100] + letterTable[100] + 3 + letterTable[((x%100)/10)*10] + letterTable[x%10]
	} else if x%100 == 0 {
		return letterTable[x/100] + letterTable[100]
	} else {
		return letterTable[x/100] + letterTable[100] + 3 + letterTable[x%100]
	}

}
