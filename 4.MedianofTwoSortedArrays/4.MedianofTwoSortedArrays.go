package main

import "fmt"

func main(){
	fmt.Println("findMedianSortedArrays: ",findMedianSortedArrays( []int{2}, []int{} ))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1, l2 := len(nums1), len(nums2)
	l := l1+l2
	v1,v2,i,j, target := 0,0,0,0,l/2

	for k:=0;k<=target;k++{
		if i<l1 && j<l2{
			if nums1[i] <= nums2[j] {
				v2=nums1[i]
				i++
			}else {
				v2=nums2[j]
				j++	
			}
		}else if i<l1{
			v2=nums1[i]	
			i++
		}else if j<l2{
			v2=nums2[j]	
			j++	
		}
		if(k==target){
			break
		}
		v1 = v2
	}
	if l%2==0{
		return float64(v1+v2)/2
	}else {
		return float64(v2)
	}
}