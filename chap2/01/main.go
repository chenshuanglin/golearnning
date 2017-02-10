package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func intsToString(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range value {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.FormatInt(int64(6), 2))
	fmt.Println(intsToString([]int{3, 4, 5}))
	x, _ := strconv.Atoi("3444")
	y, _ := strconv.ParseInt("23323", 10, 32)
	fmt.Println(x, y)
}
