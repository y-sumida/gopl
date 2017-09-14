package main

import (
	"fmt"
	"os"
)
func basename(s string) string {
	// 最後の / とその前をすべて破棄する
	for i := len(s) - 1; i >=0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 最後の '.' より前のすべてを保持する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	input := os.Args[1]
	fmt.Println(basename(input))
}
