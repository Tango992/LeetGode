package main

import "fmt"

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
}

func rearrangeArray(nums []int) []int {
	negativeNums := make([]int, 0)
	positiveNums := make([]int, 0)
	result := make([]int, 0)

	for _, num := range nums {
		if num < 0 {
			negativeNums = append(negativeNums, num)
			continue
		}
		positiveNums = append(positiveNums, num)
	}

	for i := range nums {
		if i%2 == 0 {
			result = append(result, positiveNums[0])
			positiveNums = positiveNums[1:]
			continue
		}
		result = append(result, negativeNums[0])
		negativeNums = negativeNums[1:]
	}

	return result
}
