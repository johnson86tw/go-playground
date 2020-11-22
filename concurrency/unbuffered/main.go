package main

import (
	"fmt"
	"time"
)

func main() {
	// buffered channel 可以一直丟資料進去，不會等待讀出
	c := make(chan bool, 1)

	// unbuffered channel 讀寫必須完成主程式才會停止
	// c := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done")
		<-c
	}()
	c <- true
	fmt.Println("finished")

}
