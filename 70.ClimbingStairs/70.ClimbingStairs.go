package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2)) //2
	fmt.Println(climbStairs(3)) //3
}
func climbStairs(n int) int {
	return dp(0, n, make(map[int]int))
}

func dp(curr, end int, memo map[int]int) int {
	if curr > end {
		return 0
	}
	if curr == end {
		return 1
	}
	if v, okay := memo[curr]; okay {
		return v
	}
	result := dp(curr+1, end, memo) + dp(curr+2, end, memo)
	memo[curr] = result
	return result
}
