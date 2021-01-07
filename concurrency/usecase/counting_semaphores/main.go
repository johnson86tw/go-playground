package main

import (
	"log"
	"math/rand"
	"time"
)

// Seat ...
type Seat int

// Bar ...
type Bar chan Seat

// ServeCustomer ...
func (bar Bar) ServeCustomer(customerID int, seat Seat) {
	log.Print("customer#", customerID, " enters the bar")
	log.Print("++ customer#", customerID, " drinks at seat#", seat)

	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))

	log.Print("-- customer#", customerID, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// the bar has 10 seats
	bar24x7 := make(Bar, 10)

	// Place seats in an bar
	for seatID := 0; seatID < cap(bar24x7); seatID++ {
		bar24x7 <- Seat(seatID)
	}

	// 每一秒增加一位 customer
	for customerID := 0; ; customerID++ {
		time.Sleep(time.Second)
		// bar有空位才會繼續執行下去
		seat := <-bar24x7
		go bar24x7.ServeCustomer(customerID, seat)
	}

	for {
		time.Sleep(time.Second)
	}
}
