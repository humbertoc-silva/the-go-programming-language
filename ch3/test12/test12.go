package main

import "fmt"

func main() {
	fmt.Println(btoi(true))
	fmt.Println(btoi(false))
	fmt.Println(itob(1))
	fmt.Println(itob(0))
}

// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob reports whether i is non-zero
func itob(i int) bool {
	return i != 0
}
