package main

import "fmt"

func main() {
	var ages map[string]int
	if age, ok := ages["bob"]; !ok {
		fmt.Println("bob is not a key in this map")
	} else {
		fmt.Println("bob's age ", age)
	}
}
