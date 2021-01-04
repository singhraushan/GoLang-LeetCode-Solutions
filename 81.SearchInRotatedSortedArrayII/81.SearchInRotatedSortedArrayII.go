package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 1))
}

func search(nums []int, target int) bool {

	for s, e := 0, len(nums)-1; s <= e; s++ {
		if nums[s] == target || nums[e] == target {
			return true
		}
		e--
	}
	return false
}
