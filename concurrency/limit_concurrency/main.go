package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const concurrencyProcesses = 10
	const jobCount = 9

	var wg sync.WaitGroup
	wg.Add(jobCount)

	found := make(chan int)
	limitCh := make(chan struct{}, concurrencyProcesses)

	for i := 0; i < jobCount; i++ {

		limitCh <- struct{}{}

		go func(val int) {
			defer func() {
				wg.Done()
				<-limitCh
			}()

			waitTime := rand.Int31n(1000)
			fmt.Println("job: ", val, "wait time: ", waitTime, "ms")
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
			found <- val
		}(i)
	}

	// 等全部都結束，關閉 channel
	go func() {
		wg.Wait()
		close(found)
	}()

	var res []int

	// 接收所有 found channel 回來的訊息
	for p := range found {
		fmt.Println("Finished Job", p)
		res = append(res, p)
	}

	fmt.Println("Done! result: ", res)

}
