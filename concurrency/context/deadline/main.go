package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)

	d := time.Now().Add(5 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	logger.Println("start")

	go do(ctx)

	time.Sleep(10 * time.Second)
	logger.Println("The End")
}

func do(ctx context.Context) {
	if deadline, ok := ctx.Deadline(); ok {
		logger.Println("Start doing something, deadline: ", deadline)
	}

	for {
		select {
		case <-ctx.Done():
			logger.Println(ctx.Err())
			return
		default:
			logger.Println("do something")
			time.Sleep(1 * time.Second)
		}
	}
}
