package main

import (
	"errors"
	"fmt"
	"os"
)

// the error returned by the openFile function wraps *os.Patherror error which is not catched by the dot(‘.’) assert
// but is catched by As function

func main() {
	err := openFile("a")
	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("using assert: %s\n", e)
	} else {
		fmt.Println("using assert: unwrapped error")
	}

	var pathError *os.PathError

	err = openFile("a")
	if errors.As(err, &pathError) {
		fmt.Printf("using As: %s\n", pathError)
	} else {
		fmt.Println("using As: unwrapped error")
	}
}

func openFile(filename string) error {
	_, err := os.Open("a")
	if err != nil {
		return fmt.Errorf("Error opening: %w", err)
	}
	return nil
}
