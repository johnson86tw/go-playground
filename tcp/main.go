package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

const (
	commandLength = 12
)

func main() {
	startServer()
}

func startServer() {
	ln, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	req, err := ioutil.ReadAll(conn)
	defer conn.Close()

	if err != nil {
		log.Panic(err)
	}

	command := bytesToCmd(req[:commandLength])
	fmt.Println("command: ", command)

	switch command {
	case "foo":
		fmt.Println("foo command")
	case "bar":
		fmt.Println("bar command")
	default:
		fmt.Println("Unknown command")
	}
}

func bytesToCmd(bs []byte) string {
	var cmd []byte

	for _, b := range bs {
		if b != 0x0 {
			cmd = append(cmd, b)
		}
	}

	return fmt.Sprintf("%s", cmd)
}
