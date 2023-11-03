package main

import (
	"fmt"
	// "slices"
)

func main() {
	// list := []string{"hello", "world"}
	// list = slices.Insert(list, len(list), "push")
	// list = slices.Delete(list, len(list) - 1, len(list))

	res := buildArray([]int{1,3}, 3)
	// res := buildArray([]int{1,2}, 4)
	fmt.Println(res)
}

func buildArray(target []int, n int) []string {
	var result []string

	current := 1
	
	for _, t := range target {
		for current < t {
			result = append(result, "Push", "Pop")
			current++
		}
		result = append(result, "Push")
		current++
	}
	return result
}

/* Less optimized answer

func buildArray(target []int, n int) []string {
	newValue := 1
	iterator := 0
	
	var intList []int
	var strList []string
	
	for {
		intList = append(intList, newValue)
		strList = append(strList, "Push")
		
		if newValue != target[iterator] {
			intList = slices.Delete(intList, len(intList) - 1, len(intList))
			strList = append(strList, "Pop")
		} else {
			iterator++
		}
		newValue++

		if slices.Equal(intList, target) {
			break
		}
	}
    return strList
}
*/