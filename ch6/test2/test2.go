package main

import "fmt"

type TypeA struct {
	A string
	TypeB
}

// func (t TypeA) doSomething() {
// 	fmt.Println("From A")
// }

type TypeB struct {
	B string
	TypeC
}

// func (t TypeB) doSomething() {
// 	fmt.Println("From B")
// }

type TypeC struct{ C string }

func (t TypeC) doSomething() {
	fmt.Println("From C")
}

func main() {
	var typeA TypeA
	typeA.A = "A"
	typeA.B = "B"
	typeA.C = "C"
	typeA.doSomething()
}
