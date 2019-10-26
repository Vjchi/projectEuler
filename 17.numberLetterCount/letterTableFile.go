package main

func initTable() map[int]int {
	var x = make(map[int]int)

	x[0] = 0 // zero not counted, not mentionned when in number
	x[1] = 3
	x[2] = 3
	x[3] = 5
	x[4] = 4
	x[5] = 4
	x[6] = 3
	x[7] = 5
	x[8] = 5
	x[9] = 4
	x[10] = 3
	x[11] = 6
	x[12] = 6
	x[13] = 8
	x[14] = 8
	x[15] = 7
	x[16] = 7
	x[17] = 9
	x[18] = 8
	x[19] = 8
	x[20] = 6
	x[30] = 6
	x[40] = 5
	x[50] = 5
	x[60] = 5
	x[70] = 7
	x[80] = 6
	x[90] = 6
	x[100] = 7

	return x
}
