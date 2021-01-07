package main

import "fmt"

func main() {
	//  one-capacity buffered channels can be used as one-time binary semaphore.
	//  In fact, such channels can also be used as multi-time binary semaphores, a.k.a., mutex locks,
	//  though such mutex locks are not efficient as the mutexes provided in the sync standard package.
	mutex := make(chan struct{}, 1)

	counter := 0

	// lock-through-send
	// increase1 := func() {
	// 	mutex <- struct{}{} // lock
	// 	counter++
	// 	<-mutex // unlock
	// }

	// lock-through-receive
	mutex <- struct{}{}
	increase2 := func() {
		<-mutex // lock
		counter++
		mutex <- struct{}{} // unlock
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000000; i++ {
			increase2()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})

	go increase1000(done)
	go increase1000(done)
	go increase1000(done)

	<-done
	<-done
	<-done

	fmt.Println("counter: ", counter)
	fmt.Println("End")
}
