package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		defer close(naturals)
		for x := 0; x < 10; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		defer close(squares)
		for x := range naturals {
			squares <- x * x
		}
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
