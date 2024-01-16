package main

import (
	"fmt"
	"slices"
)

func main() {
	res := findWinners([][]int{{2,3}, {1,3}, {5,4}, {6,4}})
	fmt.Println(res)
}

func findWinners(matches [][]int) [][]int {
	loseCounts := make(map[int]int)
	for _, val := range matches {
		loseCounts[val[0]] += 0
		loseCounts[val[1]]++
	}

	answer := make([][]int, 2)
	for key, val := range loseCounts {
		if val == 0 {
			answer[0] = append(answer[0], key)
		} else if val == 1 {
			answer[1] = append(answer[1], key)
		}
	}
	slices.Sort(answer[0])
	slices.Sort(answer[1])
	return answer
}