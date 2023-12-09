package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello"
	res := reverseVowels(str)
	fmt.Println(res)
}

func reverseVowels(s string) string {
	var (
		vowels = "aiueo"
		vowelsReversed, finalString string
		vowelsCounter int
	)
	
	for _, c := range s {
		if strings.ContainsAny(strings.ToLower(string(s)), vowels) {
			vowelsReversed = string(c) + vowelsReversed
			continue
		}
	}
	
	for _, c := range s {
		if strings.ContainsRune(vowels, c) {
			finalString += string(vowelsReversed[vowelsCounter])
			vowelsCounter++
			continue
		}
		finalString += string(c)
	}
	return finalString
}