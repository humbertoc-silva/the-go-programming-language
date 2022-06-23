package main

import "fmt"

func main() {
	fmt.Println(triple(4))
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
}
