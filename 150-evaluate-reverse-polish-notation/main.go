package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := evalRPN([]string{"10","6","9","3","+","-11","","/","","17","+","5","+"})
	fmt.Println(res)
}

type Stack struct {
	value int
	next  *Stack
}

func evalRPN(tokens []string) int {
	ops := map[string]func(int, int) int{
		"+": func(i, j int) int {return i+j},
		"-": func(i, j int) int {return i-j},
		"*": func(i, j int) int {return i*j},
		"/": func(i, j int) int {return i/j},
	}
	
	firstNum, _ := strconv.Atoi(tokens[0])
	stacks := &Stack{value: firstNum}
	
	for _, val := range tokens[1:] {
		num, err := strconv.Atoi(val)
		if err != nil {
			calculationValue := ops[val](stacks.next.value, stacks.value)
			stacks = &Stack{
				value: calculationValue,
				next: stacks.next.next,
			}
			continue
		}

		stacks = &Stack{
			value: num,
			next: stacks,
		}
	}
	return stacks.value
}
