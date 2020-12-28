package main

import "fmt"

func main() {

	res := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}) //[[1,6],[8,10],[15,18]]
	for _, v := range res {
		fmt.Print(v)
	}
	fmt.Println("\n------------------------")
	res = merge([][]int{{1, 4}, {4, 5}}) //[[1,5]]
	for _, v := range res {
		fmt.Print(v)
	}
	fmt.Println("\n------------------------")
	res = merge([][]int{{27, 32}, {127, 134}, {375, 380}, {72, 78}, {333, 340}, {421, 422}, {116, 121}, {457, 465}, {178, 185}, {449, 454}, {481, 485}, {391, 391}, {186, 194}, {180, 184}, {116, 124}, {345, 350}, {156, 158}, {232, 235}, {307, 314}, {250, 252}, {354, 355}, {465, 467}, {208, 212}, {482, 488}, {112, 114}, {463, 468}, {488, 494}, {335, 342}, {45, 52}, {491, 494}, {76, 82}, {360, 362}, {160, 164}, {115, 124}, {101, 106}, {311, 318}, {388, 396}, {130, 138}, {403, 407}, {335, 340}, {396, 396}, {312, 319}, {357, 357}, {264, 268}, {367, 372}, {119, 122}, {443, 452}, {27, 30}, {423, 428}, {422, 427}, {212, 217}, {332, 335}, {213, 217}, {401, 407}, {279, 287}, {483, 490}, {239, 247}, {116, 122}, {148, 148}, {202, 207}, {487, 488}, {459, 459}, {300, 308}, {385, 388}, {333, 342}, {100, 106}, {96, 101}, {351, 353}, {451, 453}, {130, 132}, {341, 347}, {462, 462}, {114, 122}, {28, 30}, {459, 459}, {277, 285}, {51, 59}, {229, 232}, {427, 433}, {278, 280}, {478, 487}, {10, 14}, {16, 17}, {236, 243}, {114, 121}, {50, 57}, {333, 335}, {239, 248}, {442, 449}, {61, 62}, {349, 357}, {200, 202}, {279, 280}, {307, 315}, {99, 104}, {170, 177}, {316, 317}, {152, 154}, {42, 42}, {300, 308}, {459, 461}, {167, 169}, {66, 72}, {443, 451}, {25, 29}, {262, 268}, {195, 197}, {478, 482}, {95, 102}})
	//[10 14][16 17][25 32][42 42][45 59][61 62][66 82][95 106][112 124][127 138][148 148][152 154][156 158][160 164][167 169][170 177][178 185][186 194][195 197][200 207][208 217][229 235][236 248][250 252][262 268][277 287][300 319][332 357][360 362][367 372][375 380][385 396][401 407][421 433][442 454][457 468][478 494]
	for _, v := range res {
		fmt.Print(v)
	}
}

func merge(intervals [][]int) [][]int {
	length := len(intervals)

	mergeSort(intervals, 0, length-1)

	var res [][]int
	res = append(res, []int{intervals[0][0], intervals[0][1]})

	for i := 1; i < length; i++ {
		lastIndx := len(res) - 1
		if intervals[i][0] > res[lastIndx][1] { //Not overalp
			res = append(res, []int{intervals[i][0], intervals[i][1]}) //add new
		} else if intervals[i][0] == res[lastIndx][1] || intervals[i][1] > res[lastIndx][1] { //overalp
			res[lastIndx][1] = intervals[i][1] //update only end
		}
	}
	return res
}

//nlogn
func mergeSort(intervals [][]int, startIndex, endIndex int) {
	if startIndex < endIndex {
		mid := (startIndex + endIndex) / 2
		mergeSort(intervals, startIndex, mid)
		mergeSort(intervals, mid+1, endIndex)
		sort(intervals, startIndex, mid, endIndex) //here mid is end of 1st half
	}
}
func sort(intervals1 [][]int, startIndex, mid, endIndex int) {
	startIndex2 := mid + 1 //this is start index of 2nd half array

	for startIndex <= mid && startIndex2 <= endIndex {
		if intervals1[startIndex][0] <= intervals1[startIndex2][0] {
			startIndex++
		} else {
			curVX, curVY, tempIndex := intervals1[startIndex2][0], intervals1[startIndex2][1], startIndex2

			//shifting by 1 forward from startIndex
			for startIndex != tempIndex {
				intervals1[tempIndex][0] = intervals1[tempIndex-1][0]
				intervals1[tempIndex][1] = intervals1[tempIndex-1][1]
				tempIndex--
			}
			intervals1[startIndex][0] = curVX
			intervals1[startIndex][1] = curVY

			startIndex++ //need to update start als because it already shifting forward
			mid++        //need to update mid also because of shifting.
			startIndex2++
		}
	}
}
