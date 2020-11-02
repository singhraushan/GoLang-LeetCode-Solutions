package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
		 Right *TreeNode 
}

func main() {
	t:= TreeNode{
		Val:1,
	     Left: nil,
		 Right: &TreeNode{
			Val:2,
			Left: &TreeNode{
				Val:3,
				Left: nil,
				Right: nil,
			 },
		 	Right: nil,
		 }, 
	}
	res := preorderTraversal(&t)
	for _,e:= range res{
		fmt.Println(e)
	}
}
func preorderTraversal(root *TreeNode) []int {
	slice := make([]int,0)
	recursive(root,&slice)
	return slice
}

func recursive(root *TreeNode, res *[]int) {
	if root==nil {
		return
	}
	*res = append(*res,(*root).Val)

	recursive((*root).Left,res)

	recursive((*root).Right,res)
}