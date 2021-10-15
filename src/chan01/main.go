package main

import (
	"log"
	"sync"
	"time"
)

const (
	workers = 3

	reset  = "\033[0m"
	yellow = "\033[33m"
	purple = "\033[35m"
	cyan   = "\033[36m"
)

var (
	done    = make(chan bool)
	buffout = make(chan int, workers)
	buffin  = make(chan int, workers)
	wg      sync.WaitGroup
)

func operate(inchan <-chan int, out chan<- int) {
	defer wg.Done()
	for input := range inchan {
		log.Printf("%s[operate]%s - will put %d\n", yellow, reset, input)
		time.Sleep(time.Second)
		out <- input
		log.Printf("%s[operate]%s - putted %d\n", yellow, reset, input)
	}
}

func outputter(outputs <-chan int) {
	for output := range outputs {
		log.Printf("%s[outputter]%s - %d\n", purple, reset, output)
	}
	done <- true
	log.Printf("%s[outputter]%s - putted done\n", purple, reset)
}

func main() {
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go operate(buffin, buffout)
	}
	go outputter(buffout)

	str := 0
	for j := 0; j < (workers * 3); j++ {
		str++
		log.Printf("%s[main]%s - will seed %d\n", cyan, reset, str)
		buffin <- str
		log.Printf("%s[main]%s - seeded %d\n", cyan, reset, str)
	}
	close(buffin)
	log.Printf("%s[main]%s - closed buffin\n", cyan, reset)
	wg.Wait()
	close(buffout)
	log.Printf("%s[main]%s - closed buffout\n", cyan, reset)
	<-done
}
