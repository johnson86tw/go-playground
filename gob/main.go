package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Getter ...
type Getter interface {
	Get() string
}

// Person ...
type Person struct {
	Name string
	Age  int
}

// Get ...
func (p Person) Get() string {
	return p.Name
}

func main() {
	// target
	ken := Person{"Ken", 22}
	fmt.Printf("%+v\n", ken)

	// Register concrete type
	gob.Register(Person{})

	// Encode interface
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	interfaceEncode(encoder, ken)
	fmt.Println(buf.String())

	// Decode interface
	var kenny Getter
	err := gob.NewDecoder(buf).Decode(&kenny)
	handle(err)
	name := kenny.Get()
	fmt.Println(name)

}

func interfaceEncode(encoder *gob.Encoder, g Getter) {
	err := encoder.Encode(&g)
	handle(err)
}

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
