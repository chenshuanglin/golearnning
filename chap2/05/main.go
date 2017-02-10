package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edge := graph[from]
	if edge == nil {
		edge = make(map[string]bool)
		graph[from] = edge
	}

	edge[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("dong", "xi")
	addEdge("nan", "ben")
	fmt.Println(hasEdge("chen", "shu"))
	fmt.Println(graph)
	a := []int{1, 3, 3}
	fmt.Println(a[:0])
}
