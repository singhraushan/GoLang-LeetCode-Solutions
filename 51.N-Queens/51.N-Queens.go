package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solveNQueens(4)) //[[..Q. Q... ...Q .Q..] [.Q.. ...Q Q... ..Q.]]
}

func solveNQueens(n int) [][]string {
	//step:1- create matrix and fill with .
	var matrix = make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = "."
		}
	}
	var res [][]string
	//call dfs with begin curCol=0
	dfs(matrix, 0, &n, &res)
	return res
}

//step:2- create dfs func
func dfs(matrix [][]string, colIndex int, len *int, totalRes *[][]string) {
	//base condition if colIndex already travelled last column then increment count and return.
	if *len == colIndex {
		var str []string
		for r := 0; r < *len; r++ {
			str = append(str, strings.Join(matrix[r], ""))
		}
		*totalRes = append(*totalRes, str)
		return
	}
	//Traverse from 1st row to last row.
	for row := 0; row < *len; row++ {
		if isValid(matrix, row, colIndex, *len) { //Checking [row][colIndex] is valid box to fill Q or not
			matrix[row][colIndex] = "Q"
			dfs(matrix, colIndex+1, len, totalRes)
			matrix[row][colIndex] = "." //reverse Q to back . for next Traverse.
		}
	}
}

func isValid(matrix [][]string, curRow, curCol, len int) bool {
	for r := 0; r < len; r++ { //Traverse from 1st row to last row.
		for c := 0; c < curCol; c++ { //col should iterate till curCol NOT len.
			if matrix[r][c] == "Q" && (r == curRow || r+c == curCol+curRow || r+curCol == c+curRow) { //checking where can't fill(same row ,column & diagonally) new Q.
				return false
			}
		}
	}
	return true
}
