package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

//setup methods and func to include the sort method
type stringSlice []string

func (names stringSlice) Swap(i, j int)      { names[i], names[j] = names[j], names[i] }
func (names stringSlice) Len() int           { return len(names) }
func (names stringSlice) Less(i, j int) bool { return names[i] < names[j] }

func main() {
	//first, read the file provided p022_names.txt
	byteNames, err := ioutil.ReadFile("D:/Documents/GoWorplace/src/projectEuler/22.namesScores/p022_names.txt")
	if err != nil { //log if error
		log.Fatal(err)
	}
	//then include them all in a string slice, separated by commas
	var strNames string = string(byteNames)                    //convert raw byte in string
	var sliStrNames stringSlice = strings.Split(strNames, ",") //split by comma
	//remove the untyped caracter  '"' (quotation mark)
	for i := range sliStrNames {
		sliStrNames[i] = strings.Trim(sliStrNames[i], "\"")
	}
	//And sort the names
	sort.Sort(sliStrNames)

	count := 0 //set counter for the result
	//loop over the string slice, count alphabetical value and multiply by position
	for i := range sliStrNames {
		s := []rune(sliStrNames[i]) //convert each caracters in runes slice
		var a int
		for j := range s {
			a += int(s[j]) - 64 //convert each rune of the names in their ASCII value, minus 64
		}
		a *= i + 1 //multiply the alphabetical value by the position of name (index + 1)
		count += a //add to the counter
	}
	fmt.Println(count)
	return
}
