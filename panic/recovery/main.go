package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	a := []string{"a", "b"}
	s := checkAndPrintWithRecover(a, 2)

	// When the panic is recovered
	// then the return value of a panicking function will be the default value
	// of the return types of the panicking function
	fmt.Printf("Existing Normally, output: %s\n", s)
}

func checkAndPrintWithRecover(a []string, index int) string {
	defer handleOutOfBounds()
	return checkAndPrint(a, index)
}

func checkAndPrint(a []string, index int) string {
	if index > len(a)-1 {
		panic("Out of bound")
	}

	return a[index]
}

func handleOutOfBounds() {
	// Here if r is nil then panic did not happened.
	// So if there is no panic then call to recover will return nil
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic")
		fmt.Printf("r: %v\n", r)
		fmt.Println("Stack Trace: ")
		debug.PrintStack()
	}
}
