package main

import (
	"fmt"
	"io"
	"log"
)

type myStringData struct {
	str       string
	readIndex int //default 0
}

func main() {

	// 	src := strings.NewReader("Hello Amazing World!")

	// data, _ := ioutil.ReadAll(src)
	// fmt.Printf("Read data of length %d: %s\n", len(data), data)

	src := myStringData{"Hello Amazing World!", 0}
	buffer := make([]byte, 3)

	// stream
	for {
		n, err := src.Read(buffer)
		fmt.Printf("%d bytes read, data %s\n", n, buffer[:n])
		fmt.Printf("Source: %s\n", src.str)

		if err == io.EOF {
			fmt.Println("---end_of_file---")
			fmt.Println("length of buffer is", len(buffer))
			break
		} else if err != nil {
			log.Fatalln(err)
			break
		}
	}

	fmt.Println("length of buffer is always", len(buffer))
}

// implement io.Read function for io.Reader Interface
func (m *myStringData) Read(p []byte) (n int, err error) {
	strBytes := []byte(m.str)

	if m.readIndex >= len(strBytes) {
		return 0, io.EOF //return 0 byte read
	}

	nextReadLimit := m.readIndex + len(p)

	if nextReadLimit >= len(strBytes) {
		nextReadLimit = len(strBytes)
		err = io.EOF
	}

	readBytes := strBytes[m.readIndex:nextReadLimit]
	n = len(readBytes)

	// 在此修改 buffer，因為 slice 是 Reference Type
	copy(p, readBytes)

	m.readIndex = nextReadLimit

	return
}
