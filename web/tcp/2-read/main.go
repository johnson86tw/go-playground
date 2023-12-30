package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// 設定connection的時間，10秒後自動關閉，注意！關閉的是connection，listener仍在。
	// err := conn.SetDeadline(time.Now().Add(20 * time.Second))
	// if err != nil {
	// 	log.Fatalln("connection timeout")
	// }

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		// 掃描傳入的訊息
		fmt.Println(ln)
		// 回覆訊息
		fmt.Fprintf(conn, "I heared you say %v\n", ln)
	}
	defer conn.Close()

	// connection結束後才會執行下列程式碼
	fmt.Println("code got here.")
}
