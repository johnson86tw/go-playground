package main

import "fmt"

type notPositiveError struct {
	num int
}

func (n notPositiveError) Error() string {
	return fmt.Sprintf("checkPositive: Given number %d is not a positive number", n.num)
}

type notEvenError struct {
	num int
}

func (e notEvenError) Error() string {
	return fmt.Sprintf("checkEven: Given number %d is not an even number", e.num)
}

func checkPositive(num int) error {
	if num < 0 {
		return notPositiveError{num: num}
	}
	return nil
}

func checkEven(num int) error {
	if num%2 == 1 {
		return notEvenError{num: num}
	}
	return nil
}

func checkPositiveAndEven(num int) error {
	err := checkPositive(num)
	if err != nil {
		return fmt.Errorf("checkPositiveAndEven: %w", err)
	}

	err = checkEven(num)
	if err != nil {
		return fmt.Errorf("checkPositiveAndEven: %w", err)
	}

	return nil
}

func main() {
	err := checkPositiveAndEven(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Given number is positive and even")
	}
}
