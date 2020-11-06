package main

import "fmt"

func main(){
	fmt.Println("maxProfit: ",maxProfit([]int{7,1,5,3,6,4}))
}
func maxProfit(prices []int) int {
	max,len := 0, len(prices)
	for i:=0;i<len;i++ {
		
		startPointer, endPointer, v := i+1, len-1,prices[i]
		
		for startPointer<=endPointer {
			c1 := prices[startPointer]
			if c1>v {
				if c1-v > max {
					max =  c1-v
				}
			}
			c2:=prices[endPointer]
			if c2>v {
				if c2-v > max {
					max =  c2-v
				}
			}
			startPointer +=1
			endPointer -=1
		}
	}
	return max
}