package main

import (
	"fmt"
	"math/rand"
	"time"
)

//? https://towardsdatascience.com/evolution-of-a-salesman-a-complete-genetic-algorithm-tutorial-for-python-6fe5d2b3ca35
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%v , %v\n", rand.Intn(100), rand.Intn(100))
}
