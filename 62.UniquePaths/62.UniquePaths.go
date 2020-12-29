package main

import (
	"fmt"
)

func main() {
	fmt.Println("count:", uniquePaths(3, 7))  //28
	fmt.Println("count:", uniquePaths(3, 2))  //3
	fmt.Println("count:", uniquePaths(7, 3))  //28
	fmt.Println("count:", uniquePaths(3, 3))  //6
	fmt.Println("count:", uniquePaths(51, 9)) //1916797311
}
func uniquePaths(m int, n int) int {
	return dfs(0, 0, n-1, m-1, make(map[string]int))
}
func dfs(x, y, fX, fY int, memo map[string]int) int {
	if fX == x && fY == y {
		return 1
	}
	if x > fX || y > fY {
		return 0
	}
	if v, okay := memo[fmt.Sprintf("%d-%d", x, y)]; okay {
		return v
	}
	res := dfs(x+1, y, fX, fY, memo) + dfs(x, y+1, fX, fY, memo)
	memo[fmt.Sprintf("%d-%d", x, y)] = res
	return res
}
