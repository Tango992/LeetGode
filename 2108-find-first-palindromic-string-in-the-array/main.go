package main

import "fmt"

func main() {
	res := firstPalindrome([]string{"pp"})
	fmt.Println(res)
}

func firstPalindrome(words []string) string {
	for _, word := range words {
		valid := true
		for i := range word {
			if word[i] != word[len(word)-1-i] {
				valid = false
				break
			}
		}
		if valid {return word}
	}
	return ""
}