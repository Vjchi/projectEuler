package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	fmt.Println("Hello World")
	byteNames, err := ioutil.ReadFile("D:/Documents/GoWorplace/src/projectEuler/22.namesScores/p022_names.txt")
	if err != nil { //log if error
		log.Fatal(err)
	}

	var strNames string = string(byteNames)                 //convert raw byte in string
	var sliStrNames []string = strings.Split(strNames, ",") //split by comma

	for i := range sliStrNames {
		sliStrNames[i] = strings.Trim(sliStrNames[i], "\"")
	}

	fmt.Println(sliStrNames)

	return
}
