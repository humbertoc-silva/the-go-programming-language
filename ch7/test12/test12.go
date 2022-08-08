package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok {
		fmt.Printf("%T\n", f)
	}
	if b, ok := w.(*bytes.Buffer); ok {
		fmt.Printf("%T\n", b)
	} else {
		fmt.Printf("failure: %v\n", b)
	}
}
