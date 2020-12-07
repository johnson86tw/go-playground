package main

import (
	"fmt"
	"log"

	serial "github.com/bugst/go-serial"
)

var headled = []byte("$610C1F00000027")

func main() {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyUSB_mp1713c", mode)
	if err != nil {
		log.Fatal("Fail to open USB:", err)
	}

	fmt.Println("write:", headled)
	n, err := port.Write(headled)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent %v bytes\n", n)

	buff := make([]byte, 100)

	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("Received: %v", string(buff[:n]))

	}
}
