package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int32{5, 1, 6, 2, 8, 3}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	fmt.Println(arr)
}