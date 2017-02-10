package main

import (
	"fmt"
	"log"
	"time"
)

func sum(vars ...int) int {
	defer trace("sum")()
	sum := 0
	for _, val := range vars {
		sum += val
	}
	time.Sleep(10 * time.Second)
	return sum
}

func trace(name string) func() {
	enter := time.Now()
	log.Printf("enter %s func, time: %v\n", name, enter)
	return func() {
		log.Printf("go out %s func , execute time is %s", name, time.Since(enter))
	}
}

func main() {
	fmt.Println(sum(1, 3, 4))
	fmt.Println(sum(3, 3, 4, 4, 5, 5))
}
