package main

import "fmt"

func main() {
	res := firstPalindrome([]string{"pp"})
	fmt.Println(res)
}

func firstPalindrome(words []string) string {
	for _, word := range words {
		valid := true
		for i, j := 0, len(word)-1; i <= j; i, j = i+1, j-1 {
			if word[i] != word[j] {
				valid = false
				break
			}
		}
		if valid {return word}
	}
	return ""
}