package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string
	if i, err := fmt.Scanf("%s", &name); err != nil {
		handleError(err)
	} else {
		fmt.Println(i)
		fmt.Printf("Hi %s\n", name)
	}
	if x, err := strconv.Atoi("123"); err != nil {
		handleError(err)
	} else {
		fmt.Println(x) // x is an int
	}
	if y, err := strconv.ParseInt("123", 10, 64); err != nil {
		handleError(err)
	} else {
		fmt.Println(y) // base 10, up to 64 bits
	}
}

func handleError(err error) {
	fmt.Println(err)
}
