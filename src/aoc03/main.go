package main

import (
	"fmt"
	"strconv"
)

func main() {

	gamma, epsilon, ox, co := "", "", "", ""
	for pos := 0; pos < len(inputs[0]); pos++ {
		zero := 0
		for _, bajt := range inputs {
			bit := bajt[pos : pos+1]
			if bit == "0" {
				zero++
			}
		}
		if zero > len(inputs)-zero { //? in this case 0 is most common bit
			gamma += "0"
			epsilon += "1"
			ox += "0"
			co += "1"
		} else {
			gamma += "1"
			epsilon += "0"
			ox += "1"
			co += "0"
		}
	}
	gd, _ := strconv.ParseInt(gamma, 2, 64)
	ed, _ := strconv.ParseInt(epsilon, 2, 64)
	od, _ := strconv.ParseInt(ox, 2, 64)
	cd, _ := strconv.ParseInt(co, 2, 64)
	fmt.Printf("gamma: %d\nepsilon: %d\nnum: %d\n", gd, ed, gd*ed)
	fmt.Printf("ox: %d\nco: %d\nnum: %d\n", od, cd, od*cd)
}
