package main

import (
	"log"

	"github.com/tarm/serial"
)

// mppark black headled
var x = []byte("$610C1F00000027")

func main() {
	conf := &serial.Config{Name: "/dev/ttyUSB_mp1713c", Baud: 115200}
	s, err := serial.OpenPort(conf)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		err := s.Close()
		if err != nil {
			log.Fatal("Fail to close serial port:", err)
		}
		log.Println("serial port closed")
	}()

	_, err = s.Write(x)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Write:", y)
}
