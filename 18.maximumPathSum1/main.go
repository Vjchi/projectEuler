package main

import (
	"fmt"
)

func main() {

	var intTr [15][15]int = createIntTriangle() //create the data matrix
	var sumIntTr [15][15]int                    //create the sum matrix

	for i := 0; i < 15; i++ { //copy the last row of data in sum matrix
		sumIntTr[14][i] = intTr[14][i]
	}

	var limHor int = len(intTr) - 1 //Initialize Horizontal location vars

	for i := 14; i > 0; i-- { //for all rows, starting at bottom

		for j := 0; j < limHor; j++ { //for all columns
			//if the sum of the LEFT pair > RIGHT pair
			if sumIntTr[i][j]+intTr[i-1][j] > sumIntTr[i][j+1]+intTr[i-1][j] {
				//Only keep LEFT PAIR in the sum Matrix
				sumIntTr[i-1][j] = sumIntTr[i][j] + intTr[i-1][j]
			} else { //Else keep the RIGHT pair
				sumIntTr[i-1][j] = sumIntTr[i][j+1] + intTr[i-1][j]
			}
		}
		limHor-- //Update horizontal limit
	}
	fmt.Println(sumIntTr[0][0]) //Print top number, which will be the Max possible
	return
}
