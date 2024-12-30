package main

import (
	"fmt"
	"time"
)

// go run -race race/main.go

func main() {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	fmt.Println("a is ", a)

	time.Sleep(2 * time.Second)
}
