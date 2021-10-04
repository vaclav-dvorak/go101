package main

import (
	"fmt"
	"time"
)

var (
	done    = make(chan bool)
	buffout = make(chan int)
)

func operate(in <-chan string, out chan<- int) {
	defer close(buffout) //? this is only possible because we run one worker
	for input := range in {
		fmt.Printf("will put %d\n", len(input))
		time.Sleep(time.Second)
		out <- len(input)
		fmt.Printf("putted %d\n", len(input))
	}
	fmt.Println("closed buffout")
}

func outputter(outputs <-chan int) {
	for output := range outputs {
		fmt.Println(output)
	}
	done <- true
	fmt.Println("putted done")
}

func main() {
	buffin := make(chan string)

	go operate(buffin, buffout)
	go outputter(buffout)

	str := ""
	for j := 0; j < 4; j++ {
		str += "a"
		fmt.Printf("will seed %d\n", len(str))
		buffin <- str
		fmt.Printf("seeded %d\n", len(str))
	}
	close(buffin)
	fmt.Println("closed buffin")
	<-done
}
