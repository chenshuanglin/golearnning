package main

import (
	"fmt"
	"sync"
)

/*
*整体上这是一种惰性求值方式
 */

//将一组整数，转换为channel
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

func merge(outs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(in <-chan int) {
		for num := range in {
			out <- num
		}
		wg.Done()
	}

	wg.Add(len(outs))
	for _, c := range outs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	out := gen(1, 2, 3, 4, 5, 6, 67)
	sq1 := sq(out)
	sq2 := sq(out)

	for num := range merge(sq1, sq2) {
		fmt.Println(num)
	}
}
