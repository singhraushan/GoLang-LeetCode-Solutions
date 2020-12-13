package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	reorderList(&head)
	curr := &head
	for curr != nil {
		fmt.Print((*curr).Val, "->")
		curr = (*curr).Next
	}
}
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	var slc []*ListNode

	for slow != nil {
		if fast == nil {
			slc = append(slc, slow)
		} else {
			if (*fast).Next == nil || (*fast).Next.Next == nil {
				fast = nil
			} else {
				fast = (*fast).Next.Next
			}
		}
		slow = (*slow).Next
	}

	curr := head
	for i := len(slc) - 1; i >= 0; i-- {
		temp := (*curr).Next
		(*curr).Next = slc[i]
		(*curr).Next.Next = temp
		curr = temp
	}
	(*curr).Next = nil
}
