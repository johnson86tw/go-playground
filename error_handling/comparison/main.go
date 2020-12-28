package main

import (
	"errors"
	"fmt"
)

// ErrorOne ...
type ErrorOne struct{}

func (e ErrorOne) Error() string {
	return "Error One happened"
}

func main() {
	e1 := ErrorOne{}
	e2 := do()

	// 無法分辨到 wrapping error
	if e1 == e2 {
		fmt.Println("Equality Operator: Both errors are equal")
	}

	// e2 屬於 e1
	if errors.Is(e2, e1) {
		fmt.Println("e2 is e1. Is Operator: Both errors are equal")
	}

	// e1 不屬於 e2
	if errors.Is(e1, e2) {
		fmt.Println("e1 is e2. Is Operator: Both errors are equal")
	}
}

func do() error {
	return fmt.Errorf("Error 2: %w", ErrorOne{})
}
