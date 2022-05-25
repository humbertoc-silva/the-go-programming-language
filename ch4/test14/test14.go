package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)

	ages["alice"] = 32
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)

	fmt.Println(ages["bob"])
	ages["bob"] = ages["bob"] + 1 // happy birthday!
	fmt.Println(ages["bob"])
	ages["bob"] += 1
	ages["bob"]++
	fmt.Println(ages["bob"])

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
