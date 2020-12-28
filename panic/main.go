package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	go g()

	for {
		fmt.Println("listening")
		time.Sleep(1 * time.Second)
	}
}

func g() {
	fmt.Println("routine 1")
	time.Sleep(5 * time.Second)
	defer fmt.Println("routine 1 defer")

}
