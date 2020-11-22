package main

import (
	"errors"
	"fmt"
	"time"
)

// APIError ...
type APIError struct {
	code int
	msg  string
	when time.Time
}

// 不用 *APIError
func (a APIError) Error() string {
	return fmt.Sprintf("%v [%d] %s", a.when, a.code, a.msg)
}

// Car is an interface implement Drive
type Car interface {
	Drive() string
}

// Tesla ...
type Tesla struct{}

// Drive ...
func (t *Tesla) Drive() string {
	return "Run Truman, Run!"
}

// Toyota ...
type Toyota struct{}

// Drive ...
func (t *Toyota) Drive() string {
	return "BooBooBoo"
}

func main() {
	if _, err := checkSthExist("somebody"); err != nil {
		_, isAPIError := err.(APIError)
		fmt.Println(isAPIError, err.Error())
	}

	if _, err := fetchAPI(); err != nil {
		// 第一個回傳值就是自己
		e, isAPIError := err.(APIError)
		fmt.Println(isAPIError, e)
	}

	// type checking for interface
	t := createTesla()

	if !isTesla(t) {
		fmt.Println("what the f*ck")
	} else {
		fmt.Println("yes, it's Tesla!")
	}

	g := createToyota()
	if !isTesla(g) {
		fmt.Println("yes, it's not Tesla!")
	} else {
		fmt.Println("what the f*ck")
	}
}

func isTesla(c Car) bool {
	_, ok := c.(*Tesla)
	return ok
}

func createTesla() Car {
	return &Tesla{}
}

func createToyota() Car {
	return &Toyota{}
}

func fetchAPI() (string, error) {
	return "", APIError{
		code: 1002,
		msg:  "something went wrong",
		when: time.Now(),
	}
}

func checkSthExist(sth string) (bool, error) {
	if sth == "something" {
		return true, nil
	}

	return false, errors.New("something doesn't exist")
}
