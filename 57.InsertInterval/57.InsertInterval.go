package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))                            //[[2 5] [6 9]]
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) //[[1,2],[3,10],[12,16]]
	fmt.Println(insert([][]int{}, []int{5, 7}))                                          //[[5,7]]
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}))                                    //[[1,5]]
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}))                                    //[[1,7]]
	fmt.Println(insert([][]int{{1, 5}}, []int{5, 7}))
}
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}

	length := len(intervals)
	if length == 0 {
		return [][]int{{newInterval[0], newInterval[1]}}
	}

	var res [][]int
	start, end, isDone := newInterval[0], newInterval[1], false

	for i := 0; i < length; i++ {
		lastIndx := len(res) - 1
		if isDone { //already added
			if intervals[i][0] > res[lastIndx][1] { //Not overalp
				res = append(res, []int{intervals[i][0], intervals[i][1]}) //add new
			} else if intervals[i][0] == res[lastIndx][1] || intervals[i][1] > res[lastIndx][1] { //overalp
				res[lastIndx][1] = intervals[i][1] //update only end
			}
		} else {
			if end < intervals[i][0] {
				res = append(res, []int{start, end})
				res = append(res, []int{intervals[i][0], intervals[i][1]})
				isDone = true
			} else if end == intervals[i][0] {
				res = append(res, []int{start, intervals[i][1]})
				isDone = true
			} else if start == intervals[i][1] {
				res = append(res, []int{intervals[i][0], end})
				isDone = true
			} else if start == intervals[i][0] {
				if end >= intervals[i][1] {
					res = append(res, []int{start, end})
				} else {
					res = append(res, []int{start, intervals[i][1]})
				}
				isDone = true
			} else if intervals[i][0] < start && end <= intervals[i][1] {
				res = append(res, []int{intervals[i][0], intervals[i][1]})
				isDone = true
			} else if start < intervals[i][0] {
				if end < intervals[i][1] {
					res = append(res, []int{start, intervals[i][1]})
				} else {
					res = append(res, []int{start, end})
				}
				isDone = true
			} else if start < intervals[i][1] && end > intervals[i][1] {
				res = append(res, []int{intervals[i][0], end})
				isDone = true
			} else {
				res = append(res, []int{intervals[i][0], intervals[i][1]})
			}
		}

	}
	if isDone == false {
		res = append(res, []int{start, end})
	}
	return res
}
