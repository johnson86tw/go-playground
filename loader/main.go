package main

import (
	"fmt"
	"time"
)

// Loader ...
type Loader struct {
	isLoading bool
}

func main() {
	l := NewLoader()

	go func() {
		time.Sleep(3 * time.Second)
		l.isLoading = false
	}()

	l.loading()

}

// NewLoader ...
func NewLoader() *Loader {
	return &Loader{isLoading: false}
}

func (l *Loader) loading() {
	l.isLoading = true

	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}

		if l.isLoading == false {
			fmt.Printf("\n")
			break
		}
	}
}
