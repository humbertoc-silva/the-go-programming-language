package main

import "fmt"

func main() {
	a := make([]int, 3)
	b := make([]int, 3, 5)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	fmt.Println(len(b))
	fmt.Println(cap(b))
}
