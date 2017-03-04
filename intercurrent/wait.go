package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("test,z要关闭了")
		time.Sleep(time.Second)
		close(exit)
	}()

	<-exit
	fmt.Println("OK exit")
}
