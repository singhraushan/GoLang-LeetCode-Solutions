package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := ListNode{Val: 3}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 0}
	node4 := ListNode{Val: 4}
	node.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2

	fmt.Println(detectCycle(&node))
}
func detectCycle(head *ListNode) *ListNode {
	mp := make(map[*ListNode]int)
	for head != nil {
		if _, okay := mp[head]; okay {
			return head
		}
		mp[head] = 0
		head = (*head).Next
	}
	return nil
}
