package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no/such/file")
	fmt.Println(os.IsNotExist(err)) // "true"
	fmt.Println(err)                // "open no/such/file: no such file or directory"
	fmt.Printf("%#v\n", err)

}
