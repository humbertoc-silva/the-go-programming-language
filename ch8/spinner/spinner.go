package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	start := time.Now()
	fibN := fib(n) // slow
	fmt.Printf("\nFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("fib took %v\n", time.Since(start))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
