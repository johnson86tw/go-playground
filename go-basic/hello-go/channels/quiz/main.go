package main

import (
	"fmt"
)

func main() {
	greeting := "Hi There!"

	c := make(chan string)

	go func(g string) {
		c <- "hi"
		fmt.Println(g)
	}(greeting)

	fmt.Println(<-c)

}
