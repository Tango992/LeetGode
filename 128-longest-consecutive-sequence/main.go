package main

import "fmt"

func main() {
	res := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	hashSet := make(map[int]bool, 0)
	for _, num := range nums {
		hashSet[num] = true
	}

	counter := 0
	for num := range hashSet {
		if !hashSet[num-1] {
			iteration := num + 1
			counterTmp := 1

			for hashSet[iteration] {
				iteration++
				counterTmp++
			}

			if counterTmp > counter {
				counter = counterTmp
			}
		}
	}
	return counter
}