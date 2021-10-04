package main

import (
	"fmt"
)

func shuffle(s string) string {
	var r, even, odd string

	for i, l := range s {
		if i%2 == 0 {
			even += string(l)
		} else {
			odd += string(l)
		}
	}
	r = even + " " + odd
	return r
}

func main() {
	var (
		T      int
		S      string
		output []string
	)

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&S)
		output = append(output, shuffle(S))
	}

	for _, line := range output {
		fmt.Println(line)
	}

}
