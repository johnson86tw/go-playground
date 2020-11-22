package main

import (
	"fmt"
	"io"
	"os"
)

type sampleStore struct {
	data []byte
}

func main() {
	ss := sampleStore{}
	bytesWritten, err := ss.Write([]byte("Hello"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten, err)
	fmt.Println(string(ss.data))

	// 使用 io.WriteString
	bytesWritten, err = io.WriteString(&ss, "Amazing")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten, err)
	fmt.Println(string(ss.data))

	bytesWritten, err = ss.Write([]byte("World!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten, err)
	fmt.Println(string(ss.data))

	os.Stdout.Write([]byte("Hello World\n"))
}

//  In reality, io.EOF error does not make any sense here
// because we are writing to something,
// hence you should create your own custom error.

func (ss *sampleStore) Write(p []byte) (n int, err error) {
	// 存有的資料量等於10時
	if len(ss.data) == 10 {
		return 0, io.EOF
	}

	// 可以寫入的空間
	remainingCap := 10 - len(ss.data)
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
