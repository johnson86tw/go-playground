package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// 建立 buffer
	buf := bytes.NewBufferString("Hello ")
	fmt.Println(buf.String())

	// 寫入到 buffer
	_, err := buf.WriteString("World")
	handle(err)

	fmt.Println(buf.String())

	file, err := os.Create("text.txt")
	handle(err)

	// 從 buffer 寫出
	buf.WriteTo(file)

	buf = bytes.NewBufferString("Hello World")

	// 讀出 buffer
	container := make([]byte, 3)
	buf.Read(container)
	fmt.Println(buf.String())
	fmt.Println(string(container))

	// 用 ReadBytes 讀到分隔符為止
	buf = bytes.NewBufferString("HelloGolang")
	var d byte = 'o'
	b, _ := buf.ReadBytes(d)
	fmt.Println(b)            // Hello
	fmt.Println(buf.String()) // Golang

	// 用 ReadByte 讀出第一個 byte
	bb, _ := buf.ReadByte()
	fmt.Println(string(bb)) // G

	// 用 ReadFrom 讀入 buffer
	file, _ = os.Open("text.txt")
	defer file.Close()

	var content bytes.Buffer
	content.ReadFrom(file)

	fmt.Println(content.String()) // Hello World

	// 用 Next 從 buffer 讀出前 n 個
	b = content.Next(2)
	fmt.Println(string(b)) // He
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
