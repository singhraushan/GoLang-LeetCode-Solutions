package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	/*node := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	res := reverseBetween(&node, 2, 4)*/
	node := ListNode{Val: 3}
	node2 := ListNode{Val: 5}
	node.Next = &node2
	res := reverseBetween(&node, 1, 2)
	curr := res
	for curr != nil {
		fmt.Print((*curr).Val, "->")
		next := (*curr).Next
		curr = next
	}
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var curr, prev *ListNode = head, nil
	currPos := 1
	for currPos != m {
		next := (*curr).Next
		prev = curr
		curr = next
		currPos++
	}
	if prev != nil {
		(*prev).Next = reverse(curr, n-m)
	} else {
		return reverse(curr, n-m)
	}
	return head
}

func reverse(currHead *ListNode, len int) *ListNode {
	var start, curr, prev *ListNode = currHead, currHead, nil
	for len > -1 {
		next := (*curr).Next
		(*curr).Next = prev
		prev = curr
		curr = next
		len--
	}
	(*start).Next = curr
	return prev
}
