package main

import "fmt"

func main() {
	case1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(case1))
}

func longestCommonPrefix(strs []string) string {
	tmp := strs[0]
	for _, str := range strs {
		i := 0

		for ; i < len(tmp) && i < len(str) && tmp[i] == str[i]; i++ {
			fmt.Println("i:", i, string(tmp[i]), string(str[i]))
		}
		tmp = tmp[:i]
	}
	return tmp
}
