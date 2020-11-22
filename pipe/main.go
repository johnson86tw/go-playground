package main

import (
	"fmt"
	"io"
)

func main() {
	src, dst := io.Pipe()

	go func() {
		dst.Write([]byte("DATA_1"))
		dst.Write([]byte("DATA_2"))
		dst.Close()
	}()

	buf := make([]byte, 6)

	bytesRead, err := src.Read(buf)
	fmt.Printf("bytes read: %d, value: %s, err:%v\n", bytesRead, buf, err)

	bytesRead, err = src.Read(buf)
	fmt.Printf("bytes read: %d, value: %s, err:%v\n", bytesRead, buf, err)

	bytesRead, err = src.Read(buf)
	fmt.Printf("bytes read: %d, value: %s, err:%v\n", bytesRead, buf, err)
}
