package main

/*
 *编写的该守护进程只能是在linux下运行
 *原理是利用exec，以及守护进程的父进程是1
 */

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const filename = "test.txt"

func main() {
	if os.Getppid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		return
	}
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	t := time.Tick(5 * time.Second)
	i := 0
	for {
		select {
		case <-t:
			i++
			fmt.Fprintf(f, "第%d次写入\n", i)
		}
	}
}
