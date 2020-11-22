package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(addByShareMemory(10))
	fmt.Println(addByShareCommunicate(10))
}

func addByShareMemory(n int) []int {
	var wg sync.WaitGroup
	var mux sync.Mutex

	res := []int{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			res = append(res, i)
			mux.Unlock()
		}(i)
	}

	wg.Wait()

	return res
}

func addByShareCommunicate(n int) []int {
	c := make(chan int, n)
	res := []int{}

	for i := 0; i < n; i++ {
		go func(c chan<- int, order int) {
			c <- order
		}(c, i)
	}

	// 讀出一一回來的 channel 裡面的值
	for i := range c {
		res = append(res, i)

		if len(res) == n {
			break
		}
	}

	close(c)
	return res

}
