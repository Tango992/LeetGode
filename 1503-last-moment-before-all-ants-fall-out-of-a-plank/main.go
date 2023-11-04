package main

import (
	"fmt"
)

func main() {
	n := 4
	left := []int{4, 3}
	right := []int{0, 1}
	fmt.Println(getLastMoment(n, left, right))
}

func getLastMoment(n int, left []int, right []int) int {
    ans := 0
    for _, pos := range right {
        ans = max(ans, n - pos)
    }
    for _, pos := range left {
        ans = max(ans, pos)
    }
    return ans
}

/*
Less optimized
func getLastMoment(n int, left []int, right []int) int {
	leftArray := make([]int, n+1)
	rightArray := make([]int, n+1)
	reference := make([]int, n+1)
	counter := 0

	for _, val := range left {
		leftArray = slices.Replace(leftArray, val, val+1, 1)
	}

	for _, val := range right {
		rightArray = slices.Replace(rightArray, val, val+1, 1)
	}

	for {
		leftArray = slices.Delete(leftArray, 0, 1)
		leftArray = slices.Insert(leftArray, len(leftArray), 0)

		rightArray = slices.Insert(rightArray, 0, 0)
		rightArray = slices.Delete(rightArray, len(rightArray) - 1, len(rightArray))

		if slices.Equal(leftArray, reference) && slices.Equal(rightArray, reference) {
			break
		}
		counter++
	}	
	return counter
}
*/