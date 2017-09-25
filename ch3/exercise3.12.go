package main

import (
	"os"
	"fmt"
)

func isAnagram(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	counts1 := make(map[rune]int)
	for _, s := range s1 {
		counts1[s]++
	}

	counts2 := make(map[rune]int)
	for _, s := range s1 {
		counts2[s]++
	}

	for key, value := range counts1 {
		if counts2[key] != value {
			return false
		}
	}

	return true
}

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Println(isAnagram(s1, s2))
}
