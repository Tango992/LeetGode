package main

import (
	"slices"
	"fmt"
)

func main() {
	arr := []int{0, 1, 14, 123, 1, 0, 23, 0, 0, 2, 5, 7, 1, 3, 7, 0, 0, 0}
	moveZeroes(arr)
}

func moveZeroes(nums []int)  {
    iteration := len(nums)
    for i := 0; i < iteration; i++ {
        if nums[i] == 0 {
            nums = slices.Delete(nums, i, i+1)
            nums = append(nums, 0)
			i--
            iteration--
        }
	}
	fmt.Println(iteration)
    fmt.Println(nums)
}