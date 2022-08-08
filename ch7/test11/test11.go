package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T\n", w)
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	fmt.Printf("%T\n", rw)
	h := w.(http.Handler)
	fmt.Printf("%T\n", h)
}
