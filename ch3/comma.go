package main

import (
	"os"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	input := os.Args[1]
	fmt.Println(comma(input))
}
