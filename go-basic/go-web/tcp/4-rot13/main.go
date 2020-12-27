package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		lc := strings.ToLower(ln)
		bs := []byte(lc)
		r := rot13(bs)
		fmt.Fprintf(conn, "%s - %s\n\n", bs, r)
	}

}

func rot13(bs []byte) string {
	r := make([]byte, len(bs))
	for i, e := range bs {
		if e <= 109 {
			r[i] = e + 13
		} else {
			r[i] = e - 13
		}
	}

	return string(r)
}
