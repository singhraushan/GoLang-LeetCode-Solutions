package main

import (
	"fmt"
)

func main(){
	//res:= longestConsecutive([]int{100,4,200,1,3,2})
	//res:= longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1})
	res:= longestConsecutive([]int{})
	fmt.Println(res)
}

func longestConsecutive(nums []int) int{
	
	if len(nums)<1 {
		return 0
	}
	
	max := 1
	m := make(map[int]bool)
	for _, val := range nums{
		m[val]=false
	}

	var prev, next,count int = 0,0,1;

	for _, v := range nums{
		
		//do for prev
		prev = v-1
		for {
			value,ok := m[prev]
			if ok && value==false {// means already there & not checked
				m[prev]=true
				count += 1
				prev -= 1 
			}else {
				break;
			}
		}

		//do for next
		next = v+1
		for {
			value,ok := m[next]
			if ok && value==false {
				m[next]=true
				count += 1
				next += 1 
			}else {
				break;
			}
		}
		if count>max{
			max = count
		}
		count = 1
	}

	return max
}