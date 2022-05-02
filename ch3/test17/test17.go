package main

import "fmt"

func main() {
	for i, r := range "Hello, 😂" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
