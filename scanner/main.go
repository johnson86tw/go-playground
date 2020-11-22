package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	one()
}

func basic() {
	input := "bar foo   baz"

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	// bufio.ScanLines 用來讀取 \n 的字串進行分行
	// 此外還有 ScanBytes, ScanRunes

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 1. Give me more data
func one() {
	input := "abcdefghijkl"

	scanner := bufio.NewScanner(strings.NewReader(input))

	// custom split func
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)

		if atEOF {
			return 0, nil, errors.New("bad luck")
		}
		// 當 buf 滿了就會雙倍 buf 的空間再讀一次
		// return 0, nil, nil

		// advance 2 即每次跑下兩個 byte
		return 2, nil, nil
	}

	scanner.Split(split)
	buf := make([]byte, 2)

	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}

	fmt.Println(string(buf))
}

// 2. Token found
func two() {
	input := "foofoofoo"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if bytes.Equal(data[:3], []byte{'f', 'o', 'o'}) {
			return 3, []byte{'F'}, nil
		}

		if atEOF {
			return 0, nil, io.EOF
		}

		return 0, nil, nil
	}
	scanner.Split(split)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 3. Error return then scanner stop, except ErrFinalToken
func three() {
	input := "foo end bar"
	scanner := bufio.NewScanner(strings.NewReader(input))

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil && bytes.Equal(token, []byte{'e', 'n', 'd'}) {
			return 0, []byte{'E', 'N', 'D'}, bufio.ErrFinalToken // errors.New("bad luck")
		}
		return
	}

	scanner.Split(split)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// io.EOF or bufio.ErrFinalToken 的 scanner.Err 都是 nil，因此這兩者並非真正的錯誤
	if scanner.Err() != nil {
		fmt.Printf("Error: %s\n", scanner.Err())
	}

}
