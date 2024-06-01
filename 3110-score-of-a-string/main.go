package main

import "fmt"

func main() {
	str := "hello"
	fmt.Println(scoreOfString(str))
}

func scoreOfString(s string) int {
    var counter int
    for i := range s[:len(s)-1] {
		currentCalculation :=  int(s[i]) - int(s[i+1])
        counter += currentCalculation
    }
    return counter
}
