package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func Merge(ins ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	output := func(in <-chan int) {
		for num := range in {
			out <- num
		}
		wg.Done()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for _, in := range ins {
		go output(in)
		wg.Add(1)
	}

	return out
}

func main() {
	out := gen(1, 2, 3, 4, 5, 6)
	sq1 := sq(out)
	sq2 := sq(out)

	for num := range Merge(sq1, sq2) {
		fmt.Println(num)
	}
}
