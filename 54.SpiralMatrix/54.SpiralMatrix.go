package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             //[1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})) //[1 2 3 4 8 12 11 10 9 5 6 7]
	fmt.Println(spiralOrder([][]int{{1}}))                                         //[1]
	fmt.Println(spiralOrder([][]int{{1, 2, 3}}))                                   //[1 2 3]
	fmt.Println(spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}}))                      //[1 2 3]
}

func spiralOrder(matrix [][]int) []int {
	r, c, rLen, cLen := 0, 0, len(matrix), len(matrix[0])
	var res []int
	mp := make(map[string]bool)
	for {
		//Left to right
		for _, okay := mp[fmt.Sprintf("%d-%d", r, c)]; c < cLen && !okay; {
			res = append(res, matrix[r][c])
			mp[fmt.Sprintf("%d-%d", r, c)] = true
			_, okay = mp[fmt.Sprintf("%d-%d", r, c+1)]
			if c+1 < cLen && !okay {
				c++
			} else {
				if r+1 < rLen {
					r++
				}
				break
			}
		}
		//top to bottom
		for _, okay := mp[fmt.Sprintf("%d-%d", r, c)]; r < rLen && !okay; {
			res = append(res, matrix[r][c])
			mp[fmt.Sprintf("%d-%d", r, c)] = true
			_, okay = mp[fmt.Sprintf("%d-%d", r+1, c)]
			if r+1 < rLen && !okay {
				r++
			} else {
				if c-1 >= 0 {
					c--
				}
				break
			}
		}
		//Right to left
		for _, okay := mp[fmt.Sprintf("%d-%d", r, c)]; c >= 0 && !okay; {
			res = append(res, matrix[r][c])
			mp[fmt.Sprintf("%d-%d", r, c)] = true
			_, okay = mp[fmt.Sprintf("%d-%d", r, c-1)]
			if c-1 >= 0 && !okay {
				c--
			} else {
				if r-1 >= 0 {
					r--
				}
				break
			}
		}
		//Bottom to top
		for _, okay := mp[fmt.Sprintf("%d-%d", r, c)]; r >= 0 && !okay; {
			res = append(res, matrix[r][c])
			mp[fmt.Sprintf("%d-%d", r, c)] = true
			_, okay = mp[fmt.Sprintf("%d-%d", r-1, c)]
			if r-1 < rLen && !okay {
				r--
			} else {
				if c+1 < cLen {
					c++
				}
				break
			}
		}
		//break condition
		if len(mp) == rLen*cLen {
			return res
		}
	}
	return res
}
