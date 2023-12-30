package main

import (
	"bufio"
	"fmt"
	"io"
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

	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	db := make(map[string]string)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1 {
			fmt.Fprintln(conn, "invalid command")
			continue
		}

		switch fs[0] {
		case "GET":
			input := fs[1]
			val := db[input]
			if val == "" {
				fmt.Fprintln(conn, "data not found")
			} else {
				fmt.Fprintln(conn, db[input])
			}

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "expected value")
				continue
			}
			prop := fs[1]
			value := fs[2]
			db[prop] = value
			fmt.Fprintf(conn, "data saved as %v -> %v\n", prop, value)
		case "DEL":
			delete(db, fs[1])
			fmt.Fprintf(conn, "%v has been deleted\n", fs[1])
		default:
			fmt.Fprintln(conn, "invalid command")
			continue
		}
	}
}
