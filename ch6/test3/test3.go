package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

func main() {
	mapping["test"] = "test"
	fmt.Println(Lookup("test"))
}
