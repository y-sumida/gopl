package main

import (
	"os"
	"fmt"
	"bytes"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	intStart := 0

	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		intStart++
	}
	intEnd := strings.Index(s, ".")
	if intEnd == -1 {
		intEnd = len(s)
	}
	intPart := s[intStart:intEnd]

	n := len(intPart)
	i := n%3

	if i == 0 {
		i = 3
	}
	buf.WriteString(intPart[:i])
	for ; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(intPart[i:i+3])
	}
	buf.WriteString(s[intEnd:])

	return buf.String()
}

func main() {
	input := os.Args[1]
	fmt.Println(comma(input))
}
