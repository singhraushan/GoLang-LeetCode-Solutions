package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 3}
	print(insertionSortList(&head))
	fmt.Println("\n----------")
	head = ListNode{Val: -1}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 0}
	print(insertionSortList(&head))
	fmt.Println("\n----------")
}

func insertionSortList(head *ListNode) *ListNode {
	var curr, prev *ListNode = head, nil
	for curr != nil {
		if prev == nil || (*prev).Val <= (*curr).Val {
			prev = curr
			curr = (*curr).Next
		} else {
			var currCurr, curPrev *ListNode = head, nil
			for currCurr != curr {
				if (*currCurr).Val > (*curr).Val {
					if curPrev == nil { //change head also
						head = curr
					} else {
						(*curPrev).Next = curr
					}
					(*prev).Next = (*curr).Next
					(*curr).Next = currCurr
					curr = (*prev).Next
					break
				} else {
					curPrev = currCurr
					currCurr = (*currCurr).Next
				}
			}
		}
	}
	return head
}

func print(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print((*curr).Val, "->")
		curr = (*curr).Next
	}
}
