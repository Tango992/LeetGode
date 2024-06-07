package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}

func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	dictMap := make(map[string]bool)

	for _, dict := range dictionary {
		dictMap[dict] = true
	}

	for i, word := range words {
		for j := range word {
			prefix := word[:j]

			if !dictMap[prefix] {
				continue
			}

			words[i] = prefix
			break
		}
	}

	return strings.Join(words, " ")
}
