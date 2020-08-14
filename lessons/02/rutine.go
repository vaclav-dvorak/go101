package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // sleep for 0 to 2.5 seconds
	}
	wg.Done() // Notify a task is finished.
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // register two tasks.
	go sayGreetings("hi!", 10)
	go sayGreetings("hello!", 10)
	wg.Wait() // block until all tasks are finished.
}
