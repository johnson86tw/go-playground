package main

import (
	"fmt"
	"time"
)

func main() {
	// 一定要是 unbuffered channel，main routine 才會被 block 住
	done := make(chan struct{})

	go func() {
		fmt.Print("Hello")

		time.Sleep(time.Second * 2)

		// Receive a value from the done channel,
		// to unblock the second send in main goroutine
		<-done
	}()

	// Blocked here, wait for a notification
	done <- struct{}{}
	fmt.Println(" World!")
}
