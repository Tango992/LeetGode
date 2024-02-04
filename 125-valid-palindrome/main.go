package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	runes := make([]rune, 0)
	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
			runes = append(runes, c)
		} else if c >= 'A' && c <= 'Z' {
			runes = append(runes, (c + 32))
		}
	}

	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
