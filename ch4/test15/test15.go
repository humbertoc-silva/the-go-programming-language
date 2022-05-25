package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["charlie"] = 34
	ages["alice"] = 31
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
