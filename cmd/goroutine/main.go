package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/google/gops/agent"
)

func f(from string) {
	i := 0
	for {
		i++
		fmt.Println(from, ":", i)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("hello")

	for {
		time.Sleep(time.Second)
	}

}

// func main() {
// 	go openFirefox()

// 	for {
// 		fmt.Println("tick")
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("tock")
// 	}
// }

func openFirefox() error {
	cmd := exec.Command("firefox", "--new-window", "https://www.google.com")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
