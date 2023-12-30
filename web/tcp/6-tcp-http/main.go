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
			log.Fatalln(err.Error())
			// 讓listener有持續接受的可能，而不是一個request發生錯誤後
			// server就停掉了
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	space := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			mux(conn, ln)
		}

		if space > 0 && ln == "" {
			fmt.Println("headers is end.")
			break
		}

		if space == 1 {
			fmt.Println("***")
			break
		}

		if ln == "" {
			space++
			continue
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	reqln := strings.Fields(ln)
	m := reqln[0]
	u := reqln[1]

	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/applyProcess" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<div><h2>Hello World</h2></div>
	<a href="/apply">填問卷</a>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<div><h2>Form</h2></div>
	<form method="post" action="/applyProcess">
	<label for="fname">First name:</label><br>
 	<input type="text" id="fname" name="fname" value="John"><br>
 	<label for="lname">Last name:</label><br>
 	<input type="text" id="lname" name="lname" value="Doe"><br><br>
	<input type="submit" value="送出">
	</form>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<div><h2>Apply Process</h2></div>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, body)
}
