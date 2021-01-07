package main

import (
	"log"
	"time"
)

type T = struct{}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// initialization phase
	time.Sleep(time.Second * 3 / 2)

	// 1-to-N notification. In practice, using by close channels
	// ready <- T{}
	// ready <- T{}
	// ready <- T{}
	close(ready) // broadcast notifications

	// N-to-1 notification. In practice, using by sync.WaitGroup
	<-done
	<-done
	<-done

}

func worker(id int, ready <-chan T, done chan<- T) {
	// block util receiving ready
	<-ready
	log.Print("Worker#", id, " starts.")
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")

	done <- T{}
}
