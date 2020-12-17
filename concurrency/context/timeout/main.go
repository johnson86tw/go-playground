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

	// 呼叫後三秒自動 timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go doPerSecond(ctx)
	go doForSecond(ctx)

	time.Sleep(5 * time.Second)
	logger.Println("The End")
}

func doPerSecond(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Println(ctx.Err())
			return
		default:
			logger.Println("do something per second")
			time.Sleep(1 * time.Second)
		}
	}
}

func doForSecond(ctx context.Context) {
	select {
	case <-ctx.Done():
		logger.Println(ctx.Err())
		return
	default:
		time.Sleep(1 * time.Second)
		logger.Println("done one second job")
	}
}
