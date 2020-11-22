package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		body := callout()
		c <- body

	}()

	fmt.Println(<-c)

}

func createTimer() {
	timer := time.NewTicker(1 * time.Second)

	for v := range timer.C {
		fmt.Println(v)
	}
}

func callout() string {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return string(body[:15])
}
