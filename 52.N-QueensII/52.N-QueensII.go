package main

import "fmt"

func main() {
	fmt.Println(totalNQueens(4)) //2
	fmt.Println(totalNQueens(5)) //10
}

func totalNQueens(n int) int {
	//step:1- create matrix and fill with . or something
	var matrix = make([][]rune, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = '.'
		}
	}
	var res int = 0
	//call dfs with begin curCol=0
	dfs(matrix, 0, &n, &res)
	return res
}

//step:2- create dfs func
func dfs(matrix [][]rune, colIndex int, len, count *int) {
	//base condition if colIndex already travelled last column then increment count and return.
	if *len == colIndex {
		*count = *count + 1
		return
	}
	//Traverse from 1st row to last row.
	for row := 0; row < *len; row++ {
		if isValid(matrix, row, colIndex, *len) { //Checking [row][colIndex] is valid box to fill Q or not
			matrix[row][colIndex] = 'Q'
			dfs(matrix, colIndex+1, len, count)
			matrix[row][colIndex] = '.' //reverse Q to back . for next Traverse.
		}
	}
}

func isValid(matrix [][]rune, curRow, curCol, len int) bool {
	for r := 0; r < len; r++ { //Traverse from 1st row to last row.
		for c := 0; c < curCol; c++ { //col should iterate till curCol NOT len.
			if matrix[r][c] == 'Q' && (r == curRow || r+c == curCol+curRow || r+curCol == c+curRow) { //checking where can't fill(same row ,column & diagonally) new Q.
				return false
			}
		}
	}
	return true
}
