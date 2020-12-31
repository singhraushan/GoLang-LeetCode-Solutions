package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4)) //2
	fmt.Println(mySqrt(8)) //2
}
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	low, high := 0, x
	var mid int
	var sqr int

	for {
		mid = (low + high) / 2
		if mid == low {
			return mid
		}

		sqr = mid * mid
		if sqr == x {
			return mid
		}
		if sqr > x {
			high = mid
		}
		if sqr < x {
			low = mid
		}
	}

}
