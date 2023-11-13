package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

type SeatManager struct {
	seat *IntHeap
}

func Constructor(n int) SeatManager {
	h := make(IntHeap, n)
	return SeatManager{
		seat: &h,
	}
}

// func (this *SeatManager) Reserve() int {
// 	if this.seat[0] == 0{

// 	}
// 	return 0
// }

func (this *SeatManager) Unreserve(seatNumber int) {
}

/*
type SeatManager struct {
	Slot []bool
}

func Constructor(n int) SeatManager {
	return SeatManager{
		Slot: make([]bool, n),
	}
}

func (this *SeatManager) Reserve() int {
	var seatNumber int
	for i := range this.Slot {
		if !this.Slot[i] {
			this.Slot[i] = true
			seatNumber = i + 1
			break
		}
	}
	return seatNumber
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.Slot[seatNumber - 1] = false
}
/*
/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
