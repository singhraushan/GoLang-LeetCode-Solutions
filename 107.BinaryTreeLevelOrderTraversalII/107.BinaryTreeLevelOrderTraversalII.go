package main

import "fmt"

func main(){

	root := TreeNode{3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}
	res := levelOrderBottom(&root)
	fmt.Println("res:",res)
}

func levelOrderBottom(root *TreeNode) [][]int {
   return  bfs(root)
}

func bfs(root *TreeNode) [][]int{
    if root == nil{
	return nil
	}

	var res [][]int 
	var q []TreeNode
	
	qPointer :=0
	q = append(q, *root)

	for qPointer>=0{

		var innerRes []int;
		qPointerAdd :=0
		
		for ;qPointer >= 0; qPointer-- {
			
			node := q[0]//getting 1st value of queue
			q = q[1:]// removing 1st value from queue
			innerRes = append(innerRes, node.Val) // adding value in resultset
			
			if node.Left!=nil{
				qPointerAdd +=1
				q = append(q, *node.Left)
			}

			if node.Right!=nil{
				qPointerAdd +=1
				q = append(q, *node.Right)
			}
		}
		qPointer = qPointerAdd-1
		res = append(res, innerRes)
	}

	
	//swaping here
	len := len(res)
	for i:=0;i<len/2;i++{
		temp:= res[i];
		res[i]= res[len-1-i]
		res[len-1-i]= temp
	}

	return res
}

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}