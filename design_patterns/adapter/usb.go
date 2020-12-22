package adapter

// source https://golangbyexample.com/adapter-design-pattern-go/

// story
// MacBook Pro has a USB port that is square in shape and Windows have a USB port that is circular in shape.
// You as a client have a USB cable that is square in shape so it can only be inserted in the mac laptop.

import "fmt"

type computer interface {
	insertInSquarePort()
}

type mac struct{}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

// Adapter
type windowsAdapter struct {
	windowMachine *windows
}

// Adapter 有一個方法將 square port 轉成 circle port
func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

type client struct{}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}
