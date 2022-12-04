package main

import "fmt"

func main() {
	my, match, score := 0, 0, 0
	for i := 0; i < len(input); i++ {
		switch input[i][2] {
		case 'X':
			my = 1
		case 'Y':
			my = 2
		case 'Z':
			my = 3
		}
		switch input[i][0] {
		case 'A':
			switch my {
			case 1:
				match = 3
			case 2:
				match = 6
			case 3:
				match = 0
			}
		case 'B':
			switch my {
			case 1:
				match = 0
			case 2:
				match = 3
			case 3:
				match = 6
			}
		case 'C':
			switch my {
			case 1:
				match = 6
			case 2:
				match = 0
			case 3:
				match = 3
			}
		}
		score += my + match
		// fmt.Printf("%d + %d\n", my, match)

	}
	fmt.Printf("%d\n", score)
}
