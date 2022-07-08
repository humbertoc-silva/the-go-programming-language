package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p1 := Point{1, 2}
	pptr := &p1
	pptr.ScaleBy(2)
	fmt.Println(p1)

	p2 := Point{1, 2}
	(&p2).ScaleBy(2)
	fmt.Println(p2)

	p3 := Point{1, 2}
	p3.ScaleBy(2)
	fmt.Println(p3)

	pptr.Distance(q)
	(*pptr).Distance(q)

}
