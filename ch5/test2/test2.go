package main

import "fmt"

func f1(i, j, k int, s, t string)                { /* ... */ }
func f2(i int, j int, k int, s string, t string) { /* ... */ }

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func Sin(x float64) float64 // implemented in assembly language

func main() {
	fmt.Printf("%T\n", add)   // func(int, int) int
	fmt.Printf("%T\n", sub)   // func(int, int) int
	fmt.Printf("%T\n", first) // func(int, int) int
	fmt.Printf("%T\n", zero)  // func(int, int) int
}
