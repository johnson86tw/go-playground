package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	c := afterDuration(3 * time.Second)
	fmt.Println("doing")
	<-c
	fmt.Println("end")

}

func afterDuration(t time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)

	go func(c chan<- struct{}) {
		time.Sleep(t)
		c <- struct{}{}
	}(c)

	return c
}
