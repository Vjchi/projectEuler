package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	var x float64 = math.Pow(2, 1000) //calculat whatever the fuck that is

	var text string = strconv.FormatFloat(x, 'f', 0, 64) //convert in string
	var digitList []string = strings.Split(text, "")     //split digits

	var count int //set counter for result

	for i := range digitList { //prepare loop for whole range
		var temp int                         //set variable for conversion to int
		temp, _ = strconv.Atoi(digitList[i]) //convert in int
		count += temp                        //add to counter
	} //next string digit

	fmt.Println("The sum of all digits is ", count)
	return
}
