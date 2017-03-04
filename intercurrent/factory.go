package main

import (
	"fmt"
	"sync"
)

type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)

	go func() {
		defer r.Done()
		for x := range r.data {
			fmt.Println(x)
		}
	}()

	return r
}

func main() {
	r := newReceiver()
	r.data <- 3
	r.data <- 4

	close(r.data)

	r.Wait()
}
