package main

import "fmt"

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int, 0)

	for i, val := range nums {
		if mapIdx, found := tmp[target-val]; found {
			return []int{i, mapIdx}
		}
		tmp[val] = i
	}
	return nil
}