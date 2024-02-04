package main

import (
	"fmt"
	"strconv"
)

func main() {
	case1 := []string{"neet", "code", "loves", "you"}
	res1 := arrToStr(case1)
	fmt.Println(res1)

	res2 := strToArr(res1)
	fmt.Println(res2)
}

func arrToStr(strs []string) string {
	answer := ""
	
	for _, str := range strs {
		length := strconv.Itoa(len(str))
		delimiter := "#"
		answer += length + delimiter + str
	}
	return answer
}


func strToArr(s string) []string {
	answer := make([]string, 0)
	firstNumIndex := 0
	delimiterIndex := 0
	for i, c := range s {
		if c == '#' {
			delimiterIndex = i
		} else {
			continue
		}

		characterLength, _ := strconv.Atoi(s[firstNumIndex:delimiterIndex])
		characterUpperBound := delimiterIndex+1+characterLength
		answer = append(answer, s[delimiterIndex+1:characterUpperBound])
		firstNumIndex = characterUpperBound
	}
	return answer
}