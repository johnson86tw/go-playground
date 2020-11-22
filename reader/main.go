package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	src := strings.NewReader("Hello Amazing World!")
	buf := make([]byte, 14)
	bytesRead, err := io.ReadFull(src, buf)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead, buf[:bytesRead], err)

	bs, _ := ioutil.ReadAll(src)
	fmt.Println(string(bs))

	bytesRead, err = io.ReadFull(src, buf)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead, buf[:bytesRead], err)

	bytesRead, err = io.ReadFull(src, buf)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead, buf[:bytesRead], err)

	bs, _ = ioutil.ReadAll(src)
	fmt.Println(bs) //return [] -> 整個 src 已經被 io.ReadFull 清空了！

}
