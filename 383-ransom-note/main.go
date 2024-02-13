package main

import "fmt"

func main() {
	res := canConstruct("aa", "aab")
	fmt.Println(res)
}

func canConstruct(ransomNote string, magazine string) bool {
	var charCounter [26]int
	for _, char := range magazine {
		charCounter[char-'a']++
	}

	for _, char := range ransomNote {
		index := char - 'a'
		charCounter[index]--

		if charCounter[index] < 0 {
			return false
		}
	}
	return true
}
