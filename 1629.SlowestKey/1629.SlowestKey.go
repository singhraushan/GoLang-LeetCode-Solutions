package main

import "fmt"

func main() {
	fmt.Printf("slowestKey %c", slowestKey([]int{9, 29, 49, 50}, "cbcd"))
}
func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, v := releaseTimes[0], 0
	ch, l := keysPressed[0], len(releaseTimes)

	for i := 1; i < l; i++ {
		v = releaseTimes[i] - releaseTimes[i-1]
		if v > max {
			max = v
			ch = keysPressed[i]
		} else if v == max {
			if ch < keysPressed[i] {
				ch = keysPressed[i]
			}
		}
	}
	return ch
}
