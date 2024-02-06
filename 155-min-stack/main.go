package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}

type MinStack struct {
	stack *StackNode
	min   int
}

type StackNode struct {
	next    *StackNode
	data    int
	lastMin int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	if m.stack == nil {
		m.stack = &StackNode{
			data:    val,
			lastMin: val,
		}
		m.min = val
		return
	}

	currentMinimum := min(val, m.min)

	newStack := &StackNode{
		data:    val,
		lastMin: m.min,
		next:    m.stack,
	}

	m.stack = newStack
	m.min = currentMinimum
}

func (m *MinStack) Pop() {
	m.min = m.stack.lastMin
	m.stack = m.stack.next
}

func (m *MinStack) Top() int {
	return m.stack.data
}

func (m *MinStack) GetMin() int {
	return m.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
