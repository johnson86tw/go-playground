package main

import "fmt"

// IOTA provides an automated way to create a enum in Golang.

type size uint8

const (
	small = size(iota)
	medium
	large
)

const (
	a = 2
	b
	c
)

func main() {
	var m size = 1
	m.toString() // Medium

	fmt.Println(a, b, c) // 2 2 2
}

func (s size) toString() {
	switch s {
	case small:
		fmt.Println("Small")
	case medium:
		fmt.Println("Medium")
	case large:
		fmt.Println("Large")
	}
}
