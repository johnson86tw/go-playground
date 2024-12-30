package main

import (
	"fmt"
)

// go run -race race_condition/main.go
// source: https://larrylu.blog/race-condition-in-golang-c49a6e242259
// Add lock to fix race condition

func main() {
	a := 0
	times := 10000
	c := make(chan bool)

	// var m sync.Mutex

	for i := 0; i < times; i++ {
		go func() {
			// m.Lock()
			a++
			c <- true
			// m.Unlock()
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	fmt.Printf("a = %d\n", a)
}
