package main

import "fmt"

func appendInt(s []int, y ...int) []int {
	var d []int
	dlen := len(s) + len(y)
	if dlen <= cap(s) {
		d = s[:dlen]
	} else {
		dcap := dlen
		if dcap < 2*len(s) {
			dcap = 2 * len(s)
		}
		d = make([]int, dlen, dcap)
		copy(d, s)
	}
	copy(d[len(s):], y)
	return d
}

func main() {
	x := make([]int, 0, 7)
	x = appendInt(x, 3, 4, 5, 7)
	fmt.Println(x)
}
