package main

import "fmt"

func main() {
	case1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(case1))
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string, 0)

	for _, str := range strs {
		var score [26]int
		for _, char := range str {
			score[char-'a']++
		}
		hashMap[score] = append(hashMap[score], str)
	}

	tmp := make([][]string, 0)
	for _, val := range hashMap {
		tmp = append(tmp, val)
	}
	return tmp
}