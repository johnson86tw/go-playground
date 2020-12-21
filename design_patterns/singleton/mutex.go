package singleton

import (
	"fmt"
	"sync"
)

var (
	singleInstance *single
	lock           = &sync.Mutex{}
)

type single struct{}

func getInstance(i int) *single {
	// 這是預先檢查，避免過多的 lock 被運行
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		// 這是為了確保只有一個 goroutine 可以建立實例
		if singleInstance == nil {
			fmt.Println(i, "Creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println(i, "Single Instance already created-1")
		}
	} else {
		fmt.Println(i, "Single Instance already created-2")
	}

	return singleInstance
}
