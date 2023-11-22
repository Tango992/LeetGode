package main

import "fmt"

func main() {
	gain := []int{-4,-3,-2,-1,4,3,2}
	res := largestAltitude(gain)
	fmt.Println(res)
}

func largestAltitude(gain []int) int {
	var count, maxGain int

	for i := range gain {
		count += gain[i]
		maxGain = max(count, maxGain)
	}
	return maxGain
}