package main

import (
	"fmt"
)

func main() {
	// arr := []int{2, 1, 3, 5, 4, 6, 7}
	// k := 2

	// arr2 := []int{3,2,1}
	// k2 := 10

	arr3 := []int{1,11,22,33,44,55,66,77,88,99}
	k3 := 1000000000
	
	// getWinner(arr, k)
	res := getWinner(arr3, k3)
	fmt.Println(res)
}

func getWinner(arr []int, k int) int {
	if k == 1 {
		return max(arr[0], arr[1])
	}

	if k >= len(arr) {
		return maxArray(arr)
	}

	winner := arr[0]
	var winnerCounter int
	for _, val := range arr[1:] {
		if winner < val {
			winner = val
			winnerCounter = 1
		} else {
			winnerCounter++
		}

		if winnerCounter == k {
			return winner
		}
	}
	return winner
}

func maxArray(arr []int) int {
	var max int
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}