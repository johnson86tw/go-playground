package main

import (
	"bytes"
	"io"
	"net"
)

const (
	commandLength = 12
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// var bs []byte = make([]byte, commandLength)
	// var cmd []byte

	// cmdStr := "apple"
	// cmd = append(cmd, []byte(cmdStr)...)
	// cmd = append(cmd, bs[:commandLength-len(cmdStr)]...)
	buff := new(bytes.Buffer)
	buff.Grow(commandLength)
	buff.WriteString("bar")

	io.Copy(conn, buff)
}
