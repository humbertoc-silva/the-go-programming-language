package main

import (
	"fmt"
	"sync"
)

func main() {
	var x, y int

	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()

	go func() {
		y = 2                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
}
