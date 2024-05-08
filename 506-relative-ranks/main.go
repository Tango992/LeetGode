package main

import (
	"fmt"
	"sort"
)

func main() {
	result := findRelativeRanks([]int{10, 3, 8, 9, 4})
	fmt.Println(result)
}

func findRelativeRanks(score []int) []string {
	titleMap := createMap(score)

	var titleSlice []string
	for _, s := range score {
		titleSlice = append(titleSlice, titleMap[s])
	}
	return titleSlice
}

func createMap(scores []int) map[int]string {
	var copiedScores []int
	copiedScores = append(copiedScores, scores...)

	sort.Slice(copiedScores, func(i, j int) bool {
		return copiedScores[j] < copiedScores[i]
	})

	titleMap := make(map[int]string)
	for i, score := range copiedScores {
		switch i {
		case 0:
			titleMap[score] = "Gold Medal"
		case 1:
			titleMap[score] = "Silver Medal"
		case 2:
			titleMap[score] = "Bronze Medal"
		default:
			titleMap[score] = fmt.Sprintf("%d", i+1)
		}
	}
	return titleMap
}
