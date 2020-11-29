package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 做什麼：簡單印出一個有0-9數字的陣列，利用goroutine一個一個塞數值進陣列裡
// 適用場景：多工處理後，把處理完成的結果塞進同一個變數（記憶體）裡

func main() {
	// fmt.Println(addByShareMemory(10))
	fmt.Println(addByShareCommunicate(10))
	fmt.Println(runtime.NumCPU())
}

// 對同一個res進行append，但是使用mux去避免覆蓋現象
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

// 利用讀取channel來append到res
func addByShareCommunicate(n int) []int {
	c := make(chan int, n)
	res := []int{}

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(c chan<- int, order int) {
			defer wg.Done()
			c <- order
		}(c, i)
	}

	wg.Wait()
	close(c)

	// 讀出一一回來的 channel 裡面的值
	for i := range c {
		res = append(res, i)

		// 如果不加這段，這個for迴圈會一直等channel，就會噴 deadlock 錯誤
		// 也可以使用 wg.Wait 來取代這一行
		// if len(res) == n {
		// 	break
		// }
	}

	return res

}
