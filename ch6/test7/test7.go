package main

import (
	"fmt"
	"time"
)

type Rocket struct { /* ... */
}

func (r *Rocket) Launch() {
	fmt.Println("Launching the rocket!")
}

func main() {
	r := new(Rocket)
	time.AfterFunc(1*time.Second, r.Launch)
}
