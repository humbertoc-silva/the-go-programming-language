package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("a", "a.1")
	addEdge("a", "a.2")
	addEdge("b", "b.1")
	fmt.Println(graph)
	fmt.Println(hasEdge("a", "a.1"))
	fmt.Println(hasEdge("a", "a.3"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
