package main

import "fmt"

type CountBytes int

func (c *CountBytes) Write(p []byte) (int, error) {
	*c += CountBytes(len(p))
	return len(p), nil
}

func main() {
	var c CountBytes
	c.Write([]byte("chenshuanglin"))
	fmt.Println(c)

	c = 0
	name := "chenshuanglin"
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println(c)
}
