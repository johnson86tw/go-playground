package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// buffered channel 不會等待讀出
	// c := make(chan bool, 1)

	// unbuffered channel 讀寫必須完成主程式才會停止
	c := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			fmt.Println("Done")
			c <- true
		}()
	}

	wg.Wait()
	close(c)
	// time.Sleep(3 * time.Second)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("finished")

}
