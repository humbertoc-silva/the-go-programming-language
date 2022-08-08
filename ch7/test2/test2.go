package main

import (
	"bytes"
	"io"
)

func main() {
	// *bytes.Buffer must satisfy io.Writer
	var _ io.Writer = (*bytes.Buffer)(nil)
}
