package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	fmt.Println(maxSubArray([]int{-1, 1, 2, 1}))                   //4
}

//O(n)
func maxSubArray(nums []int) int {
	max, maxBySum := nums[0], nums[0]
	isFirst := true
	for _, v := range nums {
		if isFirst {
			isFirst = false
		} else {
			if maxBySum+v > v {
				maxBySum += v
			} else {
				maxBySum = v
			}
			if maxBySum > max {
				max = maxBySum
			}
		}
	}
	return max
}

//O(n^2)
func maxSubArray2(nums []int) int {
	len, max := len(nums), 0
	sum := make([]int, len)

	for index, value := range nums {
		if index == 0 {
			sum[index] = value
			max = value
		} else {
			sum[index] = sum[index-1] + value
			if value > max {
				max = value
			}
			if sum[index] > max {
				max = sum[index]
			}
		}
	}
	var currSum int
	for i := 1; i < len; i++ {
		for j := i + 1; j < len; j++ {
			currSum = sum[j] - sum[i-1]
			if currSum > max {
				max = currSum
			}
		}
	}
	return max
}
