package main

import "fmt"

func main() {
	fmt.Println(countOperationsToEmptyArray([]int{3,4,-1}))
	fmt.Println(countOperationsToEmptyArray([]int{1,2,4,3}))
	fmt.Println(countOperationsToEmptyArray([]int{1,2,3}))
}

/*
	https://leetcode.com/problems/make-array-empty
*/
func countOperationsToEmptyArray(nums []int) int64 {
	var operations int64

	for {
		var currentSmallest int
		smallest := false
		operations++

		for i := range nums {
			if nums[currentSmallest] > nums[i] {
				currentSmallest = i
			}
		}

		if currentSmallest == 0 {
			smallest = true
		}
		
		if smallest {
			nums = nums[1:]
		} else {
			firstValueBuffer := nums[0]
			nums = nums[1:]
			nums = append(nums, firstValueBuffer)
		}

		if len(nums) == 0 {
			break
		}
	}
	return operations
}
