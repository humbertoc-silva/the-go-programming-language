package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// pp := new(Point)
	// *pp = Point{1, 2}
	pp := &Point{1, 2}
	fmt.Println(pp)
}
