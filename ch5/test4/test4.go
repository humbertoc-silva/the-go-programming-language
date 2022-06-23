package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "read failed: %v", err)
		}
		fmt.Fprintln(os.Stdout, r)
	}
}
