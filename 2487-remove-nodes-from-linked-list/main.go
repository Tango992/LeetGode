package main

import "fmt"

// Represents a singly linked list.
type ListNode struct {
	// The linked list value.
	Val int
	// The next linked list node.
	Next *ListNode
}

func main() {
	fifthNode := &ListNode{8, nil}
	fourthNode := &ListNode{3, fifthNode}
	thirdNode := &ListNode{13, fourthNode}
	secondNode := &ListNode{2, thirdNode}
	firstNode := &ListNode{5, secondNode}

	result := removeNodes(firstNode)
	iterateNodes(result)
}

// Iterate ListNode and prints it to the stdout.
func iterateNodes(head *ListNode) {
	for {
		fmt.Print(head.Val)
		head = head.Next

		if head == nil {
			break
		}

		fmt.Print(" -> ")
	}
	fmt.Println()
}

func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	next := removeNodes(head.Next)
	if head.Val < next.Val {
		return next
	}

	head.Next = next
	return head
}
