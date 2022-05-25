package main

import "fmt"

func main() {
	var stack []int
	stack = push(stack, 1)
	stack = push(stack, 2)
	fmt.Println(top(stack))
	stack = pop(stack)
	fmt.Println(top(stack))
	stack = pop(stack)
	fmt.Println(top(stack))
}

func push(stack []int, v int) []int {
	return append(stack, v) // push v
}

func pop(stack []int) []int {
	return stack[:len(stack)-1] // pop
}

func top(stack []int) int {
	if len(stack) == 0 {
		return -1
	}
	return stack[len(stack)-1] // top of stack
}
