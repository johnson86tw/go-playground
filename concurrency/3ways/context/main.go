package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx, "Johnson")
	fmt.Println("client release connection, time to notify A and B exit")

	time.Sleep(5 * time.Second)

	fmt.Println("Ok, stop your jobs, everyone!")
	cancel()

	// 確保工作結束
	time.Sleep(1 * time.Second)

}

func foo(ctx context.Context, name string) {
	go bar(ctx, name)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "A Exit")
			return
		case <-time.After(1 * time.Second):
			fmt.Println(name, "A is doing something per second")
		}
	}
}

func bar(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "B Exit")
			return
		case <-time.After(2 * time.Second):
			fmt.Println(name, "B is doing something twice per second")
		}
	}
}
