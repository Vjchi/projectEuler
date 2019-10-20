package main

import (
	"fmt"
	"strconv"
	"strings"

	"projectEuler/13.largeSum/pkg1"
)

func main() {

	var numbStringList []string = strings.Split(pkg1.RawNumb, "\n")
	var count, countA, countB, countC, countD, countE int //create our different
	//counters to work
	for i := 0; i < len(numbStringList); i++ {
		runeNumb := []rune(numbStringList[i])
		A, _ := strconv.Atoi(string(runeNumb[0:10])) //A deals with 10 first digits
		B, _ := strconv.Atoi(string(runeNumb[10:20]))
		C, _ := strconv.Atoi(string(runeNumb[20:30]))
		D, _ := strconv.Atoi(string(runeNumb[30:40]))
		E, _ := strconv.Atoi(string(runeNumb[40:])) //E deals with last 10 digits
		countA += A                                 //countA is counting everything above 10^40
		countB += B                                 //countB is counting everything above 10^30
		countC += C                                 //countC is counting everything above 10^20
		countD += D                                 //countD is counting everything above 10^10
		countE += E                                 //countA is counting everything unit level to 10^10
	}

	countD += countE / 10000000000 //then we report any overflow from our counters
	countC += countD / 10000000000 //stuff like "carry the one" and so on
	countB += countC / 10000000000
	countA += countB / 10000000000 //finishing off with biggest power

	trim := []rune(strconv.Itoa(countA))  //change to string
	trim = trim[0:10]                     //select 10 first characters
	count, _ = strconv.Atoi(string(trim)) //revert back to String for result

	fmt.Println("The large sum is ", count) //Print result
	return
}
