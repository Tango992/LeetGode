package main

import "fmt"

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		var answer string

		if i%3 == 0 {
			answer += "Fizz"
		}

		if i%5 == 0 {
			answer += "Buzz"
		}

		if len(answer) < 1 {
			answer =  fmt.Sprintf("%v", i)
		}

		fmt.Println(answer)
	}
}