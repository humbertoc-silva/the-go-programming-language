package main

import "fmt"

func main() {
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}

func f(...int) {}
func g([]int)  {}
