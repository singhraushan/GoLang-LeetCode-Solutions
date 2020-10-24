package main

import ("fmt"
"math"
)

func main(){
	fmt.Printf("max int32: %d\n", math.MaxInt32)
	fmt.Printf("min int32: %d\n", math.MinInt32)
	fmt.Printf("Reverse value is %d.", reverse(1534236469))
}
func reverse(x int) int {
	r, min,max := 0, math.MinInt32, math.MaxInt32
	for x!=0 {
	 r = r*10 + x%10
	 x /= 10
	}
	if r>max || r<min {
		return 0
	}
    return r
}