package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

/*
	https://leetcode.com/problems/find-mode-in-binary-search-tree/
*/
func main() {
	tree := generateTree()
	result := findMode(tree)
	fmt.Println(result)
}

func findMode(root *TreeNode) []int {
	resChan := make(chan int, 10000)

	walkTheTree(root, resChan)
	close(resChan)

	maps := map[int]int{}
	var maxRecurring int

	for val := range resChan {
		maps[val]++

		if maps[val] > maxRecurring {
			maxRecurring++
		}
	}

	var result []int
	for key, val := range maps {
		if val == maxRecurring {
			result = append(result, key)
		}
	}
	return result
}

func walkTheTree(tree *TreeNode, ch chan<- int) {
	if tree == nil {
		return
	}
	ch <- tree.Val
	walkTheTree(tree.Left, ch)
	walkTheTree(tree.Right, ch)
}

func generateTree() *TreeNode {
	tree := &TreeNode{Val: 1}

	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 5}

	tree.Left.Left = &TreeNode{Val: 3}
	tree.Left.Right = &TreeNode{Val: 4}

	tree.Right.Left = &TreeNode{Val: 3}
	tree.Right.Right = &TreeNode{Val: 5}

	tree.Left.Left.Left = &TreeNode{Val: 2}
	tree.Left.Left.Right = &TreeNode{Val: 2}

	tree.Right.Left.Left = &TreeNode{Val: 5}

	return tree
}