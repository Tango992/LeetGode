package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree
*/
func main() {
	tree := &TreeNode{Val: 4}

	tree.Left = &TreeNode{Val: 8}
	tree.Left.Left = &TreeNode{Val: 0}
	tree.Left.Right = &TreeNode{Val: 1}

	tree.Right = &TreeNode{Val: 5}
	tree.Right.Right = &TreeNode{Val: 6}


	res := averageOfSubtree(tree)
	fmt.Println(res)
}

func averageOfSubtree(root *TreeNode) int {
	var equalNodeCount int

	calculateAvg(root, &equalNodeCount)

	return equalNodeCount
}

func calculateAvg(root *TreeNode, equalNodeCount *int) {
	var sum, count int
	
	if root == nil {
		return
	}

	calculateAvg(root.Left, equalNodeCount)

	inOrderTraversal(root, &sum, &count)
	avg := sum / count
	if avg == root.Val {
		*equalNodeCount++
	}

	calculateAvg(root.Right, equalNodeCount)
}


func inOrderTraversal(root *TreeNode, sum, count *int) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left, sum, count)
	
	*sum += root.Val
	*count++
	
	inOrderTraversal(root.Right, sum, count)
}