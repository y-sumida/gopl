package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string][]string)

	inputFiles := os.Args[1:]
	if len(inputFiles) == 0 {
		countLines(os.Stdin, counts, files)
	} else {
		for _, arg := range inputFiles {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, files[line])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, files map[string][]string) {
	filename := f.Name()
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !contains(files[text], filename) {
			files[text] = append(files[text], filename)
		}
	}
}

func contains(array []string, value string) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}
