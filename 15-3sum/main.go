package main

import (
	"fmt"
	"slices"
)

func main() {
	// cases := []int{-1, 0, 1, 2, -1, -4}
	// cases := []int{0,0,0}
	cases := []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	fmt.Println(threeSum(cases))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	answer := make([][]int, 0)

	for i := 0; i < len(nums) && nums[i] < 1; i++ {
		iVal := nums[i]

		if i > 0 && nums[i-1] == iVal {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			target := iVal * -1
			jVal := nums[j]
			kVal := nums[k]

			if 	len(answer) > 0 && 
				answer[len(answer)-1][0] == iVal && 
				answer[len(answer)-1][1] == jVal && 
				answer[len(answer)-1][2] == kVal {
				j++
				continue
			}
			
			if jVal+kVal < target {
				j++
			} else if jVal+kVal > target {
				k--
			} else if jVal+kVal == target {
				answer = append(answer, []int{iVal, jVal, kVal})
				j++
			}
		}
	}
	return answer
}