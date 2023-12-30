package main

import (
	"fmt"
	"io"
	"os"
)

type logWrite struct{}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("can't not find the file.")
		os.Exit(1)
	}

	lw := logWrite{}
	io.Copy(lw, file)
}

func (logWrite) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return n, err
}
