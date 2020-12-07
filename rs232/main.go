package main

import (
	"fmt"
	"log"

	serial "github.com/bugst/go-serial"
)

// 00000B9124000006000000000001B9
var x = []byte{10, 11, 12, 13, 50, 51, 52, 53, 62, 63, 64, 65, 66, 67, 68, 69}
var y = []byte{0x00, 0x00, 0x0B, 0x91, 0x24, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xB9}
var z = "00000B9124000006000000000001B9"

func main() {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal("Fail to open USB:", err)
	}

	fmt.Println("write:", []byte(z))
	n, err := port.Write([]byte(z))
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
		fmt.Printf("%v", string(buff[:n]))

	}
}
