package main

import "fmt"

func NotZero() (ret int) {
	defer func() {
		if q := recover(); q != nil {
			ret = 3
		}
	}()
	panic("test")
}

func main() {
	fmt.Println(NotZero())
}
