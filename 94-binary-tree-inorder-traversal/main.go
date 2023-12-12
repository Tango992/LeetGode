package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 2}
	tree.Right.Left = &TreeNode{Val: 3}

	fmt.Println(inorderTraversal(tree))
}

func inorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	traverse(root, &arr)
	return arr
}

func traverse(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, arr)
	*arr = append(*arr, root.Val)
	traverse(root.Right, arr)
}
