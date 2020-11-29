package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second)
	count := 0

	for v := range timer.C {
		fmt.Println(count, v)
		count++
	}
}
