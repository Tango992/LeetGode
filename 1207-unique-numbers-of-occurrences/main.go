package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1,2,3,4,1,2,3,4,1,2,3,4,}))
}

func uniqueOccurrences(arr []int) bool {
	occur := map[int]int{}

	for _, val := range arr {
		occur[val]++
	}

	uniqueOccur := map[int]bool{}
	for _, val := range occur {
		uniqueOccur[val] = true
	}

	if len(occur) != len(uniqueOccur) {
		return false
	}
	return true
}