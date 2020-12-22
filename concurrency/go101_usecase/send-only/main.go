package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Pass send-only channels as arguments

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := make(chan int32), make(chan int32)

	go shortTermRequest(a)
	go longTermRequest(b)

	fmt.Println(<-a, <-b)
	fmt.Println("done")

}

func shortTermRequest(r chan<- int32) {
	time.Sleep(3 * time.Second)
	r <- rand.Int31n(100)
}

func longTermRequest(r chan<- int32) {
	time.Sleep(5 * time.Second)
	r <- rand.Int31n(100)
}
