package main

import "fmt"

func main() {
	switch s := suit(); s {
	case "Spades":
		fmt.Println("Spades")
	case "Hearts":
		fmt.Println("Hearts")
	case "Diamonds":
		fmt.Println("Diamonds")
	case "Clubs":
		fmt.Println("Clubs")
	default:
		panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
	}
}

func suit() string {
	return "Hearts"
}
