package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	i := 0
	wg.Add(3) // 3 jobs to do

	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1 done")
		i++
	}()

	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2 done")
		i++
	}()

	go func() {
		defer wg.Done()
		fmt.Println("goroutine 3 done")
		i++
	}()

	wg.Wait() //wait for task to be done

	fmt.Printf("All %v jobs done!\n", i)
}
