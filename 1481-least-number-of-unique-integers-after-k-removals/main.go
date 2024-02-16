package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	numMap := make(map[int]int)
	for _, num := range arr {
		numMap[num]++
	}

	var maxCount int
	countMap := make(map[int][]int)
	for num, count := range numMap {
		maxCount = max(maxCount, count)
		for i := 0; i < count; i++ {
			countMap[count] = append(countMap[count], num)
		}
	}

	arr = make([]int, 0)
	for i := 1; i <= maxCount; i++ {
		if _, ok := countMap[i]; !ok {
			continue
		}
		arr = append(arr, countMap[i]...)
	}

	var counter, currentNum int
	for _, num := range arr[k:] {
		if currentNum != num {
			currentNum = num
			counter++
		}
	}
	return counter
}