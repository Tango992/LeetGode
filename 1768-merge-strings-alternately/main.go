package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := "abc"
	word2 := "pqr"
	res := mergeAlternately(word1, word2)
	fmt.Println(res)
}

func mergeAlternately(word1, word2 string) string {
	maxIteration := max(len(word1),  len(word2))
	temp := []string{}
	
	for i := 0; i < maxIteration; i++ {
		if i < len(word1) {
			temp = append(temp, string(word1[i]))
		}
		if i < len(word2) {
			temp = append(temp, string(word2[i]))
		}
	}
	return strings.Join(temp, "")
}