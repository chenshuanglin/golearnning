package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	line := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		linestr := input.Text()
		if linestr == "q" {
			fmt.Println("bye")
			os.Exit(0)
		}
		if !line[linestr] {
			line[linestr] = true
			fmt.Println(linestr)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reader error : v%\n", err)
		os.Exit(1)
	}
}
