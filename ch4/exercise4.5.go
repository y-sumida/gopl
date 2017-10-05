package main

import "fmt"

func unique(s []string) []string {
	i := 0
	for _, value := range s {
		if s[i] != value {
			i++
			s[i] = value
		}
	}
	return s[:i+1]
}

func main() {
	s := []string{"one", "two", "two", "three", "two"}
	fmt.Println(unique(s))
}
