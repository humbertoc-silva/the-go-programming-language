package main

import "fmt"

var u uint8 = 225
var i int8 = 127

func main() {
	fmt.Println(u, u+1, u*u) // "255 0 1"
	fmt.Println(i, i+1, i*i) // "127 -128 1"
	fmt.Println("Humberto" > "Luana")
}
