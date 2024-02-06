package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	for low, high := 0, len(nums)-1; low <= high; {
		mid := (low + high) / 2
		midVal := nums[mid]

		if midVal == target {
			return mid
		}

		if midVal < target {
			low = mid + 1
		} else if midVal > target {
			high = mid - 1
		}
	}
	return -1
}
