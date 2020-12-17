package main

import (
	"context"
	"log"
	"os"
	"time"
)

type key string

var k key = "name"
var logger *log.Logger

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())

	c1 := context.WithValue(ctx, k, 1)
	c2 := context.WithValue(ctx, k, 2)
	c3 := context.WithValue(ctx, k, 3.1)

	go watch(c1)
	go watch(c2)
	go watch(c3)

	time.Sleep(4 * time.Second)

	logger.Println("Stop!")
	cancel()

	// 3秒才夠正確的停止 goroutine
	time.Sleep(3 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Println("Mission: ", ctx.Value(k), " stopping")
			return
		default:
			if val, ok := ctx.Value(k).(int); ok {
				logger.Println("Mission: ", val, "is working")
				time.Sleep(time.Duration(val) * time.Second)
			} else {
				logger.Println("Mission error")
				<-ctx.Done()
				return
			}
		}
	}
}
