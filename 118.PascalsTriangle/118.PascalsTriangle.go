package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}
func generate(numRows int) [][]int {
	var res [][]int
	for row := 0; row < numRows; row++ {
		var r []int
		for col := 0; col <= row; col++ {
			if col == 0 {
				r = append(r, 1)
			} else if len(res[row-1]) > col {
				r = append(r, res[row-1][col-1]+res[row-1][col])
			} else {
				r = append(r, 1)
			}
		}
		res = append(res, r)
	}
	return res
}
