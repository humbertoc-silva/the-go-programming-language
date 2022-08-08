package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // success: f == os.Stdout
	fmt.Println(f)
	c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	fmt.Println(c)
}
