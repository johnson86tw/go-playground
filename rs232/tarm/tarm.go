package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tarm/serial"
)

var x = []byte{0xAA, 0x01, 0x0f, 0x00, 0x00, 0xBA}

func main() {

	conf := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("writing... %v\n", x)
	_, err = s.Write(x)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")

	log.Println("Start Reading...")

	buf := make([]byte, 128)

	for {
		n, err := s.Read(buf)

		if err == io.EOF {
			fmt.Println("--- end of file ---")
			break
		} else if err != nil {
			log.Println("ERROR: ", err)
			break
		}

		log.Printf("Received: %s\n", buf[:n])
	}

}

func listen(s *serial.Port, c chan<- int) {
	log.Println("Start Reading...")
	buf := make([]byte, 128)

	for {
		fmt.Println("a")
		n, err := s.Read(buf)
		fmt.Println("b")
		if err == io.EOF {
			fmt.Println("--- end of file ---")
			break
		} else if err != nil {
			log.Println("ERROR: ", err)
			break
		}

		log.Println("Received:", buf[:n])
	}

	c <- 1
}

func reading(s *serial.Port, c chan<- int) {
	log.Println("Start Reading...")

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		fmt.Println("scan")
		log.Println(scanner.Text())
	}

	c <- 1
}
