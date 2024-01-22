package main

import (
	"fmt"
)

func main() {
	res := findErrorNums([]int{4,8,1,5,2,7,4,6})
	fmt.Println(res)
}

func findErrorNums(nums []int) []int {
	report := make([]int, 2)
	n := len(nums)
	count := make([]int, n)

    for _, num := range nums {
        count[num-1]++
    }

    for i:=0; i<n; i++ {
        if count[i] == 2 {
            report[0] = i+1
        }

        if count[i] == 0 {
            report[1] = i+1
        }
    }
    return report
}