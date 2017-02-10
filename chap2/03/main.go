package main

import "fmt"

func noneEmpty1(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

func noneEmpty2(s []string) []string {
	out := s[:0]
	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	src := []string{"时间", "", "年份"}
	src1 := []string{"世界", "年份", "", "真的"}
	fmt.Println(noneEmpty1(src))
	fmt.Println(src)
	fmt.Println(noneEmpty2(src1))
	fmt.Println(src1)
}
