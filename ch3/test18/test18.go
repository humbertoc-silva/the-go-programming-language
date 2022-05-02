package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for range s {
		n++
	}
	fmt.Println(n)

	fmt.Println(utf8.RuneCountInString(s))

	x := "世界"
	fmt.Printf("% x\n", x)
	r := []rune(x)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))

	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
}
