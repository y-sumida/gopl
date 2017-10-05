package main

import "fmt"

func rotate(s []int) []int {
	first := s[0]
	copy(s, s[1:])
	s[len(s)-1] = first

	return s
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(s))
}
