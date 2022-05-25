package main

import "fmt"

var m = make(map[string]int)

func main() {
	s := []string{"foo", "bar"}
	Add(s)
	Add(s)
	Add(s)
	fmt.Println(Count(s))
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
