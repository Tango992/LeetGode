package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	var waterArea int
	for i, j := 0, len(height)-1; i < j; {
		leftHeight := height[i]
		rightHeight := height[j]
		length := j - i
		tall := min(leftHeight, rightHeight)
		currentWaterArea := length * tall

		if currentWaterArea > waterArea {
			waterArea = currentWaterArea
		}

		if leftHeight < rightHeight {
			i++
			continue
		} else if leftHeight > rightHeight {
			j--
			continue
		}
		i++
	}
	return waterArea
}
