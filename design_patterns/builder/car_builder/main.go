package main

import "fmt"

// https://github.com/crazybber/awesome-patterns/blob/master/creational/builder.md

// Builder pattern separates the construction of a complex object from its representation
// so that the same construction process can create different representations.

// package car

// Speed ...
type Speed float64

const (
	// MPH ...
	MPH Speed = 1
	// KPH ...
	KPH = 1.60934
)

// Color ...
type Color string

const (
	// BlueColor ...
	BlueColor Color = "blue"
	// GreenColor ...
	GreenColor = "green"
	// RedColor ...
	RedColor = "red"
)

// Wheels ...
type Wheels string

const (
	// SportsWheels ...
	SportsWheels Wheels = "sports"
	// SteelWheels ...
	SteelWheels = "steel"
)

// Builder ...
type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

// Interface ...
type Interface interface {
	Drive() error
	Stop() error
}

// CarBuilder ...
type CarBuilder struct {
	Car
}

// Car ...
type Car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

// NewBuilder ...
func NewBuilder() Builder {
	return &CarBuilder{}
}

// Paint ...
func (cb *CarBuilder) Paint(color Color) Builder {
	cb.color = color
	return cb
}

// Wheels ...
func (cb *CarBuilder) Wheels(wheels Wheels) Builder {
	cb.wheels = wheels
	return cb
}

// TopSpeed ...
func (cb *CarBuilder) TopSpeed(speed Speed) Builder {
	cb.speed = speed
	return cb
}

// Build ...
func (cb *CarBuilder) Build() Interface {
	return &cb.Car
}

// Drive ...
func (c *Car) Drive() error {
	fmt.Printf("Car Info: %s %s %v\n", c.color, c.wheels, c.speed)
	fmt.Println("Start Driving!")
	return nil
}

// Stop ...
func (c *Car) Stop() error {
	fmt.Println("Stop!")
	return nil
}

// client
func main() {
	assembly := NewBuilder().Paint(RedColor)

	familyCar := assembly.Wheels(SteelWheels).TopSpeed(MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(SportsWheels).TopSpeed(KPH).Build()
	sportsCar.Drive()
}
