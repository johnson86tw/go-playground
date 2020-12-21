package singleton

import (
	"sync"
	"testing"
)

const count = 500

func TestParallelSingleton(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(count)

	instances := [count]*single{}

	for i := 0; i < count; i++ {
		go func(index int) {
			instances[index] = initSingleton()
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("worker instance is not equal")
		}
	}
}

func TestSingletonBasic(t *testing.T) {
	ins1 := initSingleton()
	ins2 := initSingleton()

	if ins1 != ins2 {
		t.Fatal("inistance is not exactly the same")
	}
}
