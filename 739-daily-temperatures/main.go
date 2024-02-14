package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
}


func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	for i := len(temperatures)-2; i >= 0; i-- {
		j := i+1

		for temperatures[j] <= temperatures[i] {
			if result[j] == 0 {
				break
			}
			j += result[j]
		}

		if temperatures[j] > temperatures[i] {
			result[i] = j - i
		}
	}
	return result
}

