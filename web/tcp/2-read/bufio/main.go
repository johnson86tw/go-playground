package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := "Hello. My name is Johnson.\nHow are you?"

	io.Copy(os.Stdout, strings.NewReader(s))

}

func bufIO(s string) {

	scanner := bufio.NewScanner(strings.NewReader(s))

	if scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
