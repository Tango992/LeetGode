package main

import "fmt"

func main() {
	two()
}

func one() {
	total := 0
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			total += 10
		} else {
			total -= 5
		}
	}
	fmt.Println(total)
}

func two() {
	a := 2
	b := 5
	c := 7
	for {
		if a > (2*b) {
			break
		}
		a += (c-b)
		b--
		c++
	}
	fmt.Println(c-a)
}