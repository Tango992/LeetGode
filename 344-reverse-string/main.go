package main

import "fmt"

func main() {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(a)
	fmt.Println(string(a))
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}