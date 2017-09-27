package main

import (
	"fmt"
)

const (
	KB float64 = 1000 // uint64 でもOverflow
	MB = 1000 * KB
	GB = 1000 * MB
	PB = 1000 * GB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Println(KB, MB, GB, PB, EB, ZB, YB)
}
