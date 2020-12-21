package singleton

import (
	"fmt"
	"testing"
)

func TestMutexSingleton(t *testing.T) {
	for i := 0; i < 100; i++ {
		go getInstance(i)
	}

	fmt.Scanln()
}
