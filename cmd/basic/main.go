package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("ls", "-al").Output()
	if err != nil {
		panic(err)
	}

	fmt.Print(string(out))
}
