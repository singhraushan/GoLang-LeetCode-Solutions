package main

import "fmt"

type ListNode struct {
	 Val int
     Next *ListNode
}
func main(){

	l1 := ListNode{
		Val:2,
		Next: &ListNode{
			Val:4,
			Next: &ListNode{
				Val:3,
			},
		},
	}
	l2 := ListNode{
		Val:5,
		Next: &ListNode{
			Val:6,
			Next: &ListNode{
				Val:4,
			},
		},
	}
	res := addTwoNumbers(&l1, &l2)
	for res != nil {
		fmt.Printf("(*res).Val : %d\n", (*res).Val)
		res = res.Next
	}


	/*l1 := ListNode{
		Val:2,
		Next: &ListNode{
			Val:4,
			Next: &ListNode{
				Val:9,
			},
		},
	}
	l2 := ListNode{
		Val:5,
		Next: &ListNode{
			Val:6,
			Next: &ListNode{
				Val:4,
				Next: &ListNode{
					Val:9,
				},
			},
		},
	}
	res := addTwoNumbers(&l1, &l2)
	for res != nil {
		fmt.Printf("(*res).Val : %d\n", (*res).Val)
		res = res.Next
	}*/
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result ListNode
	var current *ListNode
	var lsum,carry int

	for l1 != nil || l2 != nil {
		
		if l1 != nil{
			lsum =  (*l1).Val
			l1 = l1.Next
		}
		
		if l2 != nil{
			lsum = lsum +  (*l2).Val
			l2 = l2.Next
		}

		t := lsum + carry
		if current == nil {//1st time
			if t<10{
				result = ListNode{Val: t ,Next: nil}
				carry = 0
			}else {
				result = ListNode{Val: t%10 ,Next: nil}
				carry = t/10
			}
			current = &result	
		}else{
			if t<10{
				(*current).Next = &ListNode{Val: t,Next: nil}
				carry = 0
			}else{
				(*current).Next = &ListNode{Val: t%10,Next: nil}
				carry = t/10
			}
			current = (*current).Next
		}
		lsum = 0
	}

	if carry>0{
		(*current).Next = &ListNode{Val: carry,Next: nil}
	}
	return &result;
}