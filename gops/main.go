package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	pid := os.Getpid()
	fmt.Println(pid)

	for {
		time.Sleep(time.Second)
	}
}
