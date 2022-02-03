package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%v , %v\n", rand.Intn(100), rand.Intn(100))
	rnd()
}
