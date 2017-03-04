package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var rs [5]struct {
		id     int
		result int
	}

	for i := 0; i < len(rs); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)

			rs[i].id = i + 1
			rs[i].result = i + 100

		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n", rs)
}
