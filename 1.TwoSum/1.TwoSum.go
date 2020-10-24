package main

import "fmt"

func main(){

	//res:= twoSum([]int{3,2,4},6)
	//res:= twoSum([]int{0,5,5,0},0)
	res := twoSum([]int{2,5,5,11},10)
	//res:= twoSum([]int{2,7,11,15},9)
	//res:= twoSum([]int{3,3},6)

	if res!=nil{
		fmt.Printf("0th & 1st index value are: %d & %d.", res[0], res[1])
	}
	
}

func twoSum(nums []int, target int) []int{
	
	m := make(map[int]int)
	
	for i,ele := range nums {
		val,ok := m[target-ele]
		if ok {
			return []int{val,i};
		}else {
			m[ele]=i
		}
	}
	return nil;
}