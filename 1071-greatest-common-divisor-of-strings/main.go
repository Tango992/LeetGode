package main

import (
	"fmt"
)

func main() {
	str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	res := gcdOfStrings(str1, str2)
	fmt.Println(res)
}

func gcdOfStrings(s1, s2 string) string {
    if s1 + s2 != s2 + s1 {
        return ""
    }
    x := gcd(len(s1), len(s2))
    return s1[:x]
}

func gcd(a, b int) int {
	fmt.Println(a,b)
    for b != 0 {
        a, b = b, a % b
		fmt.Println(a,b)
    }
    return a
}