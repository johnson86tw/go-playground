package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		fmt.Println(i)
		go source(c)
	}

	r := <-c
	fmt.Println("spend time: ", time.Since(startTime).Seconds())
	fmt.Println("random number: ", r)
}

func source(c chan<- int32) {
	r, _ := rand.Int31(), rand.Intn(3)+1
	time.Sleep(6 * time.Second)
	c <- r
}
