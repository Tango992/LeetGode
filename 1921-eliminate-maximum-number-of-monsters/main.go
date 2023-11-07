package main

import (
	"fmt"
	"slices"
)

func main() {
	dist := []int{4,3,4}
	speed := []int{1,1,2}
	res := eliminateMaximum(dist, speed)
	fmt.Println(res)
}

func eliminateMaximum(dist []int, speed []int) int {
	if len(dist) == 1 {
		return 1
	}

	timeArr := make([]float32, 0)
	for i := range dist {
		time := (float32(dist[i]) / float32(speed[i]))
		timeArr = append(timeArr, time)
	}
	
	slices.Sort(timeArr)
	
	globalTime := 1
	for _, val := range timeArr[1:] {
		if float32(globalTime) >= val {
			break
		}
		globalTime++
	}
	return globalTime
}

/* draft 1
func eliminateMaximum(dist []int, speed []int) int {
	if len(dist) == 1 {
		return 1
	}

	var globalTime int
	for i := range dist {
		time := (float32(dist[i]) / float32(speed[i]))
		fmt.Println(globalTime, time)
		if globalTime == 0 {
			globalTime++
			continue
		}
		if float32(globalTime) >= time {
			break
		}
		globalTime++
	}
	return globalTime
}
*/

/* draft 2
func eliminateMaximum(dist []int, speed []int) int {
	if len(dist) == 1 {
		return 1
	}

	var globalTime int
	var eliminated int
	for i := range dist {
		time := (float32(dist[i]) / float32(speed[i]))
		fmt.Println(globalTime, time)
		if globalTime == 0 {
			eliminated++
			globalTime++
			continue
		}
		if float32(globalTime) >= time {
			globalTime++
			continue
		}
		eliminated++
		globalTime++
	}
	return globalTime
}
*/