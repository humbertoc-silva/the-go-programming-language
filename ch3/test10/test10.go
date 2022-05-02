package main

import "fmt"

func main() {
	result, ok := compute()
	fmt.Println(result)
	fmt.Println(ok)
}

func compute() (value float64, ok bool) {
	// ...
	failed := true
	result := 3.14
	if failed {
		return 0, false
	}
	return result, true
}
