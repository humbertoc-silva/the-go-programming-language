package main

import "fmt"

func main() {
	a := "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(a)
	b := "\u4e16\u754c"
	fmt.Println(b)
	c := "\U00004e16\U0000754c"
	fmt.Println(c)
	d := '\u4e16'
	fmt.Printf("%c\n", d)
	e := '\U00004e16'
	fmt.Printf("%c\n", e)
}
