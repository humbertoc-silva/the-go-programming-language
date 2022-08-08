package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		fmt.Println(buf)
	}
}

func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
