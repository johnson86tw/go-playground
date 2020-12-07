package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	ss "github.com/bugst/go-serial"
	"github.com/tarm/serial"
)

func main() {
	contents, _ := ioutil.ReadDir("/dev")
	found := false

	// 尋找符合 tty.usbserial or ttyUSB 的目錄
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") ||
			strings.Contains(f.Name(), "ttyUSB") {
			fmt.Println("/dev/" + f.Name())
			found = true
		}
	}

	if found == false {
		fmt.Println("------------------------")
		fmt.Println("not found")
	}

	fmt.Println("---------------")

	// 取得所有 serial port
	ports, err := ss.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// 嘗試 open 所有 serial port
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)

		c := &serial.Config{Name: port, Baud: 115200}
		s, err := serial.OpenPort(c)
		if err != nil {
			fmt.Println(err)
			continue
		}

		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])

	}
}
