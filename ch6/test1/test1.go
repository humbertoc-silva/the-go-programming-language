package main

import "fmt"

func main() {
	i2 := IntList{Value: 5}
	i1 := IntList{Value: 5, Tail: &i2}
	fmt.Println(i1.Sum())
}

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
