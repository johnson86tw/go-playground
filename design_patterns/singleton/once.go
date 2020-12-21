package singleton

import (
	"fmt"
	"sync"
)

var (
	singleton *single
	once      sync.Once
)

func initSingleton() *single {
	if singleton == nil {
		once.Do(func() {
			fmt.Println("Create Single Instance Now")
			singleton = &single{}
		})
	} else {
		fmt.Println("Single instance already created-2")
	}

	return singleton
}
