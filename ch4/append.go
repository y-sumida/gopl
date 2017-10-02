package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	a := []int{0, 1, 2}
	appendA := appendInt(a, 10)
	fmt.Println(appendA)

	b := make([]int, 3, 3)
	b[0], b[1], b[2] = 0, 1, 2
	appendB := appendInt(b, 10)
	fmt.Println(appendB)
}
