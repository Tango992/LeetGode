package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("tree"))
}

func frequencySort(s string) string {
	var freq [128]int
	for _, char := range s {
		freq[char]++
	}

	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if freq[b[i]] == freq[b[j]] {
			return b[i] > b[j]
		}
		return freq[b[i]] > freq[b[j]]
	})
	return string(b)
}
