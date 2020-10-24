package main
import "fmt"
func main(){
	//res:= twoSum([]int{3,2,4},6)
	//res:= twoSum([]int{0,5,5,0},0)
	//res:= twoSum([]int{2,5,5,11},10)
	//res:= twoSum([]int{2,7,11,15},9)
	res:= twoSum([]int{3,3},6)

	fmt.Println("0th index value: ",res[0])
	fmt.Println("1st index value: ",res[1])
}

func twoSum(nums []int, target int) []int{
	res :=[2]int{-1,-1};
	len := len(nums)
	if len <=1 {
		return res[:]
	}
	map1 := make(map[int]int)
	map1[nums[0]]=0
	for i:=1;i<len;i++ {
		val,ok := map1[target-nums[i]]
		if ok {
			res[0] = val
			res[1]=i
			return res[:];
		}else {
			map1[nums[i]]=i
		}
	}
	return res[:];
}