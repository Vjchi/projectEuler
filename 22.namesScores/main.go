package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type stringSlice []string

func (names stringSlice) Swap(i, j int)      { names[i], names[j] = names[j], names[i] }
func (names stringSlice) Len() int           { return len(names) }
func (names stringSlice) Less(i, j int) bool { return names[i] < names[j] }

func main() {

	fmt.Println("Hello World")
	byteNames, err := ioutil.ReadFile("D:/Documents/GoWorplace/src/projectEuler/22.namesScores/p022_names.txt")
	if err != nil { //log if error
		log.Fatal(err)
	}

	var strNames string = string(byteNames)                    //convert raw byte in string
	var sliStrNames stringSlice = strings.Split(strNames, ",") //split by comma

	for i := range sliStrNames {
		sliStrNames[i] = strings.Trim(sliStrNames[i], "\"")
	}
	sort.Sort(sliStrNames)

	count := 0

	for i := range sliStrNames {
		s := []rune(sliStrNames[i])
		var a int
		for j := range s {
			a += int(s[j]) - 64
		}
		a *= i
		count += a
	}
	fmt.Println(count)
	return
}
