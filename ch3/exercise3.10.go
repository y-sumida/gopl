package main

import (
	"os"
	"fmt"
	"bytes"
)

func comma(s string) string {
	var buf bytes.Buffer

	n := len(s)
	i := n%3

	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])
	for ; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i:i+3])
	}

	return buf.String()
}

func main() {
	input := os.Args[1]
	fmt.Println(comma(input))
}
