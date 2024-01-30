package main

import "fmt"

func main() {
	case1 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result := containsDuplicate(case1)
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]bool, 0)

	for _, val := range nums {
		tmp[val] = true
	}

	return len(tmp) != len(nums)
}