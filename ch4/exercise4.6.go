package main

import (
	"fmt"
	"unicode"
)

func removeContinuousSpaces(s []byte) []byte {
	output := []byte{}

	for i, char := range s {
		if unicode.IsSpace(rune(char)) {
			if i > 0 && unicode.IsSpace(rune(s[i-1])) {
				continue
			} else {
				output = append(output, ' ')
			}
		} else {
			output = append(output, char)
		}
	}
	return output
}

func main() {
	s := []byte("\t\t123\r\n\nabc")
	fmt.Println(string(removeContinuousSpaces(s)))
}
