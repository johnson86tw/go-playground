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

type sampleStore struct {
	data []byte
}

func main() {
	src := myStringData{"Hello Amazing World!", 0}
	buffer := make([]byte, 3)
	dst := sampleStore{}

	for {
		n, err := src.Read(buffer)
		fmt.Printf("%d bytes read, data: %s\n", n, buffer[:n])
		bytesWritten, _ := dst.Write(buffer[:n])
		fmt.Printf("Bytes written %d", bytesWritten)

		if err == io.EOF {
			fmt.Println("---end_of_file---")
			fmt.Println("length of buffer is", len(buffer))
			break
		} else if err != nil {
			log.Fatalln(err)
			break
		}
	}

	fmt.Println("source:", src.str)
	fmt.Println("destination:", string(dst.data))

	fmt.Println("source's readIndex:", src.readIndex)
	src.readIndex = 0

	// simple way
	io.Copy(&dst, &src)

	fmt.Println("source:", src.str)
	fmt.Println("destination:", string(dst.data))

}

func (ss *sampleStore) Write(p []byte) (n int, err error) {
	// 存有的資料量等於100時
	if len(ss.data) == 100 {
		return 0, io.EOF
	}
	// 可以寫入的空間
	remainingCap := 100 - len(ss.data)
	// 要寫的資料的大小
	writeLength := len(p)

	if remainingCap <= writeLength {
		writeLength = remainingCap
		err = io.EOF
	}

	ss.data = append(ss.data, p[:writeLength]...)

	n = writeLength
	return

}

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
