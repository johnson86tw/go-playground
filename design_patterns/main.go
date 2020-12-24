package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")

	defer fmt.Println("defer end")

	log.Fatal("fatal")
}
