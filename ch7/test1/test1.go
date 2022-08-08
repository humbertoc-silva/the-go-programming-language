package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Humberto Corrêa da Silva")
	var buf []byte = make([]byte, 8)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println(n)
		fmt.Println(err)
		fmt.Println(string(buf[:n]))
	}
}
