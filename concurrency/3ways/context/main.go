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
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func foo(ctx context.Context, name string) {
	go bar(ctx, name)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "A Exit")
			return
		case <-time.After(1 * time.Second):
			fmt.Println(name, "A is doing something")
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
			fmt.Println(name, "B is doing something")
		}
	}
}
