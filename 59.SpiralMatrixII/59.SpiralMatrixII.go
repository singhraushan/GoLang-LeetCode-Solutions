package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	startRow, startCol := 0, 0
	endRow, endCol := n-1, n-1
	count := 1
	for {
		//row: left to right
		for c := startCol; c <= endCol; c++ { //only col change
			grid[startRow][c] = count
			count++
		}
		startRow++

		//col: top to bottom
		for r := startRow; r <= endRow; r++ { //only row change
			grid[r][endCol] = count
			count++
		}
		endCol--

		//row: right to left
		for c := endCol; c >= startCol; c-- { //only col change
			grid[endRow][c] = count
			count++
		}
		endRow--

		//row: bottom to end
		for r := endRow; r >= startRow; r-- { //only row change
			grid[r][startCol] = count
			count++
		}
		startCol++

		if count > n*n {
			return grid
		}
	}
	return grid
}
