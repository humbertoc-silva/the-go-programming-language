package main

import "fmt"

func main() {
	ages := make(map[string]int) // mapping from strings to ints
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)
}
