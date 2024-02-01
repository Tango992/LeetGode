package main

import "fmt"

func main() {
	case1 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(case1, 2))
}

func topKFrequent(nums []int, k int) []int {
	// map for [values]counter
	tmp := make(map[int]int, 0)

	for _, num := range nums {
		tmp[num]++
	}

	// map for [counter]values
	hashMap := make(map[int][]int, 0)

	for values, counter := range tmp {
		hashMap[counter] = append(hashMap[counter], values)
	}

	answer := make([]int, 0)
	for i := len(nums); i >= 1; i-- {
		answer = append(answer, hashMap[i]...)
		if len(answer) >= k {
			break
		}
	}
	return answer[:k]
}