package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func main() {
	/*root:= TreeNode{
		Val:1,
		Left: &TreeNode{
			Val:2,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:3,
			Left: nil,
			Right: nil,
		},
	}*/
	root:= TreeNode{
		Val:4,
		Left: &TreeNode{
			Val:9,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:0,
			Left: nil,
			Right: nil,
		},
	}
	
	(*root.Left).Left = &TreeNode{
		Val:5,
		Left: nil,
		Right: nil,
	}

	(*root.Left).Right = &TreeNode{
		Val:1,
		Left: nil,
		Right: nil,
	}

	fmt.Println("sumNumbers:",sumNumbers(&root))
}
func sumNumbers(root *TreeNode) int {
	if root ==nil {
		return 0
	}
	totalSum :=0;
	dfs(root, &[]int{}, &totalSum);
	return totalSum
}

func dfs(root *TreeNode, value *[]int, totalSum *int ) {
	*value = append( *value ,(*root).Val)
	
	if isLeafNode(root) {
		sum:=0
		for _,e := range *value{
			sum = sum*10+e;
		}
		*totalSum = *totalSum + sum;
	}

	if (*root).Left!=nil {
		v := (*value)[:]
		dfs((*root).Left, &v, totalSum )
	}
	if (*root).Right!=nil {
		v := (*value)[:]
		dfs((*root).Right, &v, totalSum )
	}

}

func isLeafNode(root *TreeNode) bool {
	return (*root).Left==nil && (*root).Right==nil
}