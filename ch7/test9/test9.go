package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)

	var t error = errors.New("error test")
	fmt.Printf("%T %v\n", t, t)

	e := fmt.Errorf("error %s", "test")
	fmt.Println(e)
}
