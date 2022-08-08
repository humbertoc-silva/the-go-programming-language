package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	l := []string{"L", "E", "H"}
	sort.Sort(StringSlice(l))
	fmt.Println(l)

	names := []string{"Bob", "Alice"}
	sort.Strings(names)
	fmt.Println(names)
}
