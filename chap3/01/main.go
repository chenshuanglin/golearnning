package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ServerWaitResult(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server no respones: %s ; retryning\n", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s\n", url, timeout)
}

func main() {
	fmt.Println(ServerWaitResult("http://dddddd/dd"))
}
