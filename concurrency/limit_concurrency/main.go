package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Pattern: Worker Queue

func main() {
	const limitNum = 10
	const jobCount = 100

	var wg sync.WaitGroup
	wg.Add(jobCount)

	found := make(chan int)
	queue := make(chan int)

	// 100個工作
	go func(queue chan<- int) {
		counter := 0
		for i := 0; i < jobCount; i++ {
			// 把工作寫入 queue
			queue <- i
			counter++
		}

		// 因為 queue 是 unbuffered，所以此routine會在這裡被block住，等到所有queue都被讀出，才會執行下方的程式碼
		close(queue)
		fmt.Println("close channel queue, counter is ", counter)
	}(queue)

	// for i := 0; i < jobCount; i++ {
	// 	queue <- i
	// }
	// close(queue)

	// 只有十個 goroutine，每個 goroutine 都在搶 queue 裡頭的工作
	for i := 0; i < limitNum; i++ {
		go func(queue <-chan int, found chan<- int) {
			// 讀出 queue 的工作內容
			for val := range queue {
				defer wg.Done()
				waitTime := rand.Int31n(1000)
				fmt.Println("job: ", val, "wait time: ", waitTime, "ms")
				time.Sleep(time.Duration(waitTime) * time.Millisecond)
				found <- val
			}
		}(queue, found)
	}

	var res []int

	// 必須使用 goroutine，等到一百個工作完成以後，要 close channel found
	// 才不會在下方的讀 found 迴圈中讀完全部的 channel 然後還在等待著讀。因為 found 是 unbuffered
	go func() {
		wg.Wait()
		close(found)
	}()

	// 讀出所有 found channel 回來的訊息
	for p := range found {
		fmt.Println("Finished Job", p)
		res = append(res, p)
	}

	fmt.Printf("%d numbers of the jobs done!\n", len(res))

}
