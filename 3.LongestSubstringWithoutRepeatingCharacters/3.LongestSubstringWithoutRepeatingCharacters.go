package main

import (
	"fmt"
	//"strings"
)

func main() {
	fmt.Println("lengthOfLongestSubstring",lengthOfLongestSubstring("abba"))//2
	//fmt.Println("lengthOfLongestSubstring",lengthOfLongestSubstring("abbc"))//2
	//fmt.Println("lengthOfLongestSubstring",lengthOfLongestSubstring("pwwkew"))//3
	//fmt.Println("lengthOfLongestSubstring",lengthOfLongestSubstring("dedf"))//3
}

func lengthOfLongestSubstring(s string) int {
	
	if len(s)<1{
		return 0
	}
	max:=1
	var temp []rune
	for _,v := range s{
		j := strings.IndexRune(string(temp), v)
		if j>=0{
			if len(temp)>max{
				max = len(temp)
			}
			temp = temp[j+1:]
		}
		temp = append(temp,v)
	}

	if len(temp)>max{
		return len(temp)
	}
	return max
}