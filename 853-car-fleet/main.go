package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(carFleet(12, []int{10,8,0,5,3}, []int{2,4,1,1,3}))
}

type carInfo struct {
	pos int
	spd int
}

func carFleet(target int, position []int, speed []int) int {
	carFleets := make([]carInfo, 0)
	stack := make([]float32, 0)

	for i := range position {
		carFleets = append(carFleets, carInfo{pos: position[i], spd: speed[i]})
	}

	slices.SortStableFunc(carFleets, func(i, j carInfo) int {
		return cmp.Compare(j.pos, i.pos)
	})

	for _, car := range carFleets {
		stack = append(stack, float32(target - car.pos) / float32(car.spd))

		if len(stack) > 1 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack)
}
