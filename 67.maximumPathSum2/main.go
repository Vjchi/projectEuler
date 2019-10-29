package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	var intTr [100][100]int = createIntTriangle() //create the data matrix
	var sumIntTr [100][100]int                    //create the sum matrix
	for i := 0; i < 99; i++ {                     //copy the last row of data in sum matrix
		sumIntTr[99][i] = intTr[99][i]
	}

	var limHor int = len(intTr) - 1 //Initialize Horizontal location vars

	for i := 99; i > 0; i-- { //for all rows, starting at bottom

		for j := 0; j < limHor; j++ { //for all columns
			//if the sum of the LEFT pair > RIGHT pair
			if sumIntTr[i][j]+intTr[i-1][j] > sumIntTr[i][j+1]+intTr[i-1][j] {
				//Only keep LEFT PAIR in the sum Matrix
				sumIntTr[i-1][j] = sumIntTr[i][j] + intTr[i-1][j]
			} else { //Else keep the RIGHT pair
				sumIntTr[i-1][j] = sumIntTr[i][j+1] + intTr[i-1][j]
			} //try the next pair
		} //work up to the floor above
		limHor-- //Update horizontal limit
	}
	fmt.Println(sumIntTr[0][0]) //Print top number, which will be the Max possible
	return
}

func createIntTriangle() [100][100]int {

	rawDataByte, err := ioutil.ReadFile("D:/Documents/GoWorplace/src/projectEuler/67.maximumPathSum2/p067_triangle.txt")
	if err != nil { //log if error
		log.Fatal(err)
	}

	var rawData string = string(rawDataByte)                   //convert raw byte in string
	var rowsStringData []string = strings.Split(rawData, "\n") //split by return
	var x [100][100]int                                        //initialize final data

	for i := range rowsStringData { //for each row:
		var temp []string = strings.Split(rowsStringData[i], " ") //split by space

		for j := 0; j < len(temp); j++ { //for each number:
			var a int                    //initialise int
			a, _ = strconv.Atoi(temp[j]) //convert in int using Atoi
			x[i][j] = a                  //put int in slice
		}
	}
	return x
}
