package main

import "fmt"

func main() {
	res := firstUniqChar("loveleetcode")
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	var counter [26]int

	for _, c := range s {
		counter[c-'a']++
	}

	for i, c := range s {
		if counter[c-'a'] == 1 {
			return i
		}
	}
	return -1
}