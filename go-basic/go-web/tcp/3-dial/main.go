package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	// write to the connection
	fmt.Fprintln(conn, "I dial you.")
	// read from the connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		// 奇怪...下面這條傳不到server
		fmt.Fprintf(conn, "I hear you say %v", ln)
	}

	// ReadAll reads from conn 'until' an error or EOF and returns the data
	// 只所以得等20秒connection deadline時才能讀到server傳來的訊息的原因
	// bs, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(bs))
}
