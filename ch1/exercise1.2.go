package main

import (
	"fmt"
	"os"
)

func main() {
	for n, arg := range os.Args[0:] {
		fmt.Printf("%d\t%s\n", n, arg)
	}
}
