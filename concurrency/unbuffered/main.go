package main

import (
	"fmt"
)

func main() {
	// If the channel is unbuffered, the sender routine blocks until the receiver has received the value
	c := make(chan bool)

	go func() {
		fmt.Println("GO GO GO")
		<-c
	}()
	c <- true // main is the sender, so if c is unbuffered, the main routine block.

	fmt.Println("finished")

}
