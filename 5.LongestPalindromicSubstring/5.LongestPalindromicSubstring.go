package main

import "fmt"

func main(){
	fmt.Println("longestPalindrome:",longestPalindrome("ac"))
}

func longestPalindrome(s string) string {
	len := len(s)
	if len <1 {
		return ""
	}

	start, maxLen, left, right := 0, 1, 0, 0
	for i:=1;i<len;i++{
		//for even size palindrome
		left = i-1
		right = i
		for left>=0 && right<len && s[left]==s[right]{
			if right-left+1 > maxLen {
				start = left
				maxLen = right-left+1
			}
			left--
			right++
		}
		//for odd size palindrome
		left = i-1
		right = i+1
		for left>=0 && right<len && s[left]==s[right]{
			if right-left+1 > maxLen {
				start = left
				maxLen = right-left+1
			}
			left--
			right++ 
		}
	}
	return string(s[start:start+maxLen])
}
