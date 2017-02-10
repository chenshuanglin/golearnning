package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer PrintStack()
	f(3)
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(num int) {
	fmt.Printf("f(%d)\n", num+0/num)
	defer fmt.Printf("defer %d\n", num)
	f(num - 1)
}
