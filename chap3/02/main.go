package main

import "fmt"

func Square() func() int {
	var i int
	return func() int {
		i++
		return i * i
	}
}

func main() {
	f := Square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
