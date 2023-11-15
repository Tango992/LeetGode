package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	str := "  a good     example"
	res := reverseWords(str)
	fmt.Println(res)
}

func reverseWords(s string) string {
    words := strings.Fields(s)
    slices.Reverse(words)
    return strings.Join(words, " ")
}