package main

import (
	"errors"
	"fmt"
)

type inputError struct {
	message      string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

func main() {
	err := validate("john", "")
	if err != nil {
		// 抓出客製化的錯誤訊息
		if inputError, ok := err.(*inputError); ok {
			fmt.Println(inputError)
			fmt.Printf("Missing Field is %s\n", inputError.getMissingField())
		} else {
			fmt.Println("Undefined Error: ", err)
		}
	}
}

func validate(name, gender string) error {
	if name == "" {
		return &inputError{"Name is mandatory", "name"}
	}
	if gender == "" {
		return &inputError{"Gender is mandatory", "gender"}
	}

	if name == "john" {
		return errors.New("name cannot be john")
	}

	fmt.Println("Success")
	return nil
}
