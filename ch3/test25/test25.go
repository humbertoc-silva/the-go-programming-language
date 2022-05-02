package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Printf("%T %[1]v\n", FlagUp)
	fmt.Printf("%T %[1]v\n", FlagBroadcast)
	fmt.Printf("%T %[1]v\n", FlagLoopback)
	fmt.Printf("%T %[1]v\n", FlagPointToPoint)
	fmt.Printf("%T %[1]v\n", FlagMulticast)
}
