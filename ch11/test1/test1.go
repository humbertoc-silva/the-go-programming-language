package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	i := rand.Intn(15)
	fmt.Println(i)
}
