package main

import (
	"fmt"
	"learnApp/lib"
	"time"
)

func main() {
	var num int
	ch := make(chan int)
	go lib.WriteNum(ch)
	num = <-ch
	fmt.Println("receive num: ", num)
	defer close(ch)

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 4
	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- 10
	}()
	fmt.Println("go into to ")
	var i1, i2 int
	select {
	case i1 = <-ch1:
		fmt.Println("ch1 read , ", i1)
	case i2 = <-ch2:
		fmt.Println("ch2 read, ", i2)
	}
	defer close(ch1)
	defer close(ch2)

	return
}
