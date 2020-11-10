package main

import "fmt"

type Board struct {
	NailsNeeded int
	NailsDriven int
}

type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

type NailDriverPuller interface {
	NailDriver
	NailPuller
}

type Mallet struct{}

func (Mallet) DriveNail(nailSupply *int, b *Board) {
	*nailSupply--
	b.NailsDriven++
	fmt.Println("Mallet: pounded nail into the board.")
}

type Crowbar struct{}

func (Crowbar) PullNail(nailSupply *int, b *Board) {
	b.NailsDriven--
	*nailSupply++
	fmt.Println("Crowbar: yanked nail out of the board.")
}

type Contractor struct{}

func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

func (c Contractor) ProcessBoards(dp NailDriverPuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(dp, nailSupply, b)
		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(dp, nailSupply, b)
		}
	}
}

type Toolbox struct {
	// embeded two interface types
	NailDriver
	NailPuller

	nails int
}

func main() {
	boards := []Board{
		{NailsDriven: 3},
		{NailsDriven: 1},
		{NailsDriven: 6},

		{NailsNeeded: 6},
		{NailsNeeded: 9},
		{NailsNeeded: 4},
	}

	tb := Toolbox{
		NailDriver: Mallet{},
		NailPuller: Crowbar{},
		nails:      10,
	}

	var c Contractor
	c.ProcessBoards(&tb, &tb.nails, boards)

	displayState(&tb, boards)
}

func displayState(tb *Toolbox, boards []Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}

	fmt.Println()
}
