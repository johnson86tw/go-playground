package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const limitNum = 10
	const jobCount = 100

	var wg sync.WaitGroup
	wg.Add(limitNum)

	// 為什麼沒有 jobCount 就不行？
	// 因為queue是unbuffer，被塞入一百個值，一定要被讀出來，goroutine才會終止 no!
	// 只要在 found 加上 jobCount 就可以
	found := make(chan int, jobCount)
	queue := make(chan int)

	// 100個工作
	go func(queue chan<- int) {
		for i := 0; i < jobCount; i++ {
			// 把工作寫入 queue
			queue <- i
		}
		close(queue)
	}(queue)

	// for i := 0; i < jobCount; i++ {
	// 	queue <- i
	// }
	// close(queue)

	// 只有十個 goroutine，每個 goroutine 都再搶 queue 裡頭的工作
	for i := 0; i < limitNum; i++ {
		go func(queue <-chan int, found chan<- int) {
			defer wg.Done()
			// 讀出 queue 的工作內容
			for val := range queue {
				// defer wg.Done()
				waitTime := rand.Int31n(1000)
				fmt.Println("job: ", val, "wait time: ", waitTime, "ms")
				time.Sleep(time.Duration(waitTime) * time.Millisecond)
				found <- val
			}
		}(queue, found)
	}

	// 直到工作做完，關閉 channel
	// go func() {
	// 	wg.Wait()
	// 	close(found)
	// }()

	wg.Wait()
	close(found)

	var res []int

	// 讀出所有 found channel 回來的訊息
	for p := range found {
		fmt.Println("Finished Job", p)
		res = append(res, p)
	}

	fmt.Println("Done! result: ", res)

}
