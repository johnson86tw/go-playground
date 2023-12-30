package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://google.com",
		"http://golang.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}

	// for l := range c {
	// 	go func(link string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(link, c)
	// 	}(l)
	// }

	// 同上方程式碼，差別在於，上方的程式碼比較具有可讀性
	for {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
	return
}
