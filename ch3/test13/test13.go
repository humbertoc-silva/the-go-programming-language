package main

import "fmt"

func main() {
	s := "hello, wolrd"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	// c := s[len(s)] // panic: index out of range
	// fmt.Println(c)
	fmt.Println(s[0:5])            // "hello"
	fmt.Println(s[:5])             // "hello"
	fmt.Println(s[7:])             // "world"
	fmt.Println(s[:])              // "hello, world"
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	s = "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s) // "left foot, right foot"
	fmt.Println(t) // "left foot"
	// s[0] = 'L' // compile error: cannot assign to s[0]
}
