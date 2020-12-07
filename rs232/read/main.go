package main

import (
	"fmt"
	"log"

	"github.com/bugst/go-serial"
)

func main() {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyUSB_mp1713c", mode)
	if err != nil {
		log.Fatal("Fail to open USB:", err)
	}

	buff := make([]byte, 100)

	fmt.Println("Reading...")

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
