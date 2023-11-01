package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	https://leetcode.com/problems/add-two-numbers/
*/
func main() {
	v1 := &ListNode{Val: 3}
	v2 := &ListNode{Val: 4, Next: v1}
	v3 := &ListNode{Val: 2, Next: v2}

	v4 := &ListNode{Val: 4}
	v5 := &ListNode{Val: 6, Next: v4}
	v6 := &ListNode{Val: 5, Next: v5}

	result := addTwoNumbers(v3, v6)
	fmt.Println(result)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result

	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}

		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}

		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}

		tmp = tmp.Next
	}
	return result
}