package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) //4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) //-1
	fmt.Println(search([]int{1}, 0))                   //-1
}
func search(nums []int, target int) int {
	half := len(nums) / 2
	for i, j := 0, len(nums)-1; i <= half; i, j = i+1, j-1 {
		if nums[i] == target {
			return i
		}
		if nums[j] == target {
			return j
		}
	}
	return -1
}
