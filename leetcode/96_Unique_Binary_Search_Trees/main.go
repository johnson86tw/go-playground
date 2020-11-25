package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

}

func channelPlayground() {
	// 少這一行!! 必須建立 channel 實體
	c := make(chan []byte)

	go sendRequest(c)
	body := <-c
	fmt.Println(body)
	fmt.Println("end")
}

func sendRequest(c chan []byte) {
	resp, err := http.Get("https://www.google.com/")
	handle(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handle(err)

	c <- body

}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
