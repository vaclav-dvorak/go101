package main

import "math/rand"

const maxRnd = 20

func rndNum(n int) (int, int) {
	var a, b int
	for i := 0; i < n; i++ {
		if rand.Intn(maxRnd) < maxRnd/2 {
			a++
		} else {
			b++
		}
	}
	return a, b
}

func main() {
	var n = 100
	x, y := rndNum(n)
	print("Result: ", x, " + ", y, " = ", n, "?")
	println(x+y == n)
}
