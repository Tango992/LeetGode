package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := evalRPN2([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"})
	fmt.Println(res)
}

func evalRPN2(tokens []string) int {
	ops := map[string]func(int, int) int{
		"+": func(i, j int) int {return i+j},
		"-": func(i, j int) int {return i-j},
		"*": func(i, j int) int {return i*j},
		"/": func(i, j int) int {return i/j},
	}
	
	firstNum, _ := strconv.Atoi(tokens[0])
	stack := []int{firstNum}
	
	for _, val := range tokens[1:] {
		num, err := strconv.Atoi(val)
		if err != nil {
			calculatedValue := ops[val](stack[len(stack)-2], stack[len(stack)-1])
			stack = append(stack[:len(stack)-2], calculatedValue)
			continue
		}
		stack = append(stack, num)
	}
	return stack[0]
}


type Stack struct {
	value int
	next  *Stack
}

// Alternative using Stack struct
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
