package main

// select 是用來接收 channel 的 case

import "fmt"

// https://blog.wu-boy.com/2020/10/select-multiple-channel-in-golang/
// https://blog.wu-boy.com/2019/11/four-tips-with-select-in-golang/
func main() {
	c := make(chan int, 3)

	c <- 2
	select {
	case <-c:
		fmt.Println("random 01")
	case <-c:
		fmt.Println("random 02")
	default:
		fmt.Println("exit")
	}
}
