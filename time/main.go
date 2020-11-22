package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(1*time.Second + 500*time.Millisecond)

	end := time.Now()

	elapsed := end.Sub(start)

	fmt.Println(elapsed)

}
