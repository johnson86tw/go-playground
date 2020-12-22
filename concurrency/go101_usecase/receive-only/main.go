package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Return receive-only channels as results

// main routine 會暫停等到兩個 channel 變數都被讀取後，才會繼續執行下去
// 原本要花8秒的時間，並行後只需要花5秒

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := shortTermRequest(), longTermRequest()
	fmt.Println(<-a, <-b)
	fmt.Println("done")
}

func shortTermRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func longTermRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 5)
		r <- rand.Int31n(100)
	}()

	return r
}
