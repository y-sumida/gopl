package main

import "fmt"

func unique(s []string) []string {
	pre := ""
	result := []string{}
	for _, value := range s {
		if pre != value {
			result = append(result, value)
		}
		pre = value
	}
	return result
}

func main() {
	s := []string{"one", "two", "two", "three", "two"}
	fmt.Println(unique(s))
}
