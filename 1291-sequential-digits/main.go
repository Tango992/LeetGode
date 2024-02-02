package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sequentialDigits(1000, 13000))
}
func sequentialDigits(low int, high int) []int {
	result := make([]int, 0)

	for i := 1; i <= 9; i++ {
		num := i

		for j := i + 1; j <= 9; j++ {
			num = num*10 + j
			fmt.Println(num)

			if num >= low && num <= high {
				result = append(result, num)
				fmt.Println()
			}
		}
	}

	sort.Ints(result)
	return result
}