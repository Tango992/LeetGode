package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}

func majorityElement(nums []int) int {
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}

	var maxCount int
	countMap := make(map[int]int)
	for num, count := range numCount {
		countMap[count] = num
		maxCount = max(maxCount, count)
	}

	return countMap[maxCount]
}
