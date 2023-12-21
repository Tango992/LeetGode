package main

import "fmt"

func main() {
	four()
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

func three() {
	ch := []string{"Z", "X", "C"}
	tmp := ""

	for i := 0; i < 5; i++ {
		tmp = ch[0]
		ch[0] = ch[2]
		ch[1] = tmp
		ch[2] = ch[1]
	}
	fmt.Println(ch)
}

func four() {
	amount := []int{0,0,0,0,0,0,0,0,0,0}
	numbers := []int{1,2,3,4,2,9,3}
	for i := 0; i < len(numbers); i++ {
		amount[numbers[i]]++
	}
	fmt.Println(amount)
	fmt.Println(amount[1])
}