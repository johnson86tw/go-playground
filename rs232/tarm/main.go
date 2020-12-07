package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

var w = []byte("00000B9124000006000000000001B9")
var y = []byte{0x00, 0x00, 0x0B, 0x91, 0x24, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xB9}

func main() {

	conf := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	s, err := serial.OpenPort(conf)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		err := s.Close()
		if err != nil {
			log.Fatal("Fail to close serial port:", err)
		}
		log.Println("Close serial port done")
	}()

	_, err = s.Write(y)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Write:", y)

	buf := make([]byte, 128)

	log.Println("Start listening...")

	for {
		n, err := s.Read(buf)
		fmt.Println(n)
		if err != nil {
			log.Panic("ERROR:", err)
			break
		}

		if n == 0 {
			log.Panic("no content to read")
			break
		}

		log.Printf("%v", buf[:n])

	}

}
