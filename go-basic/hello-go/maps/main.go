package main

import "fmt"

func main() {
	basic()
	system := makeMap()

	for sy, os := range system {
		fmt.Println(sy, os)
	}
}

func basic() {
	petal := map[string]bool{
		"she love me": false,
		"I love her":  true,
	}

	petal["This is a dream"] = true
	fmt.Println(petal)
}

func makeMap() map[string]string {
	system := make(map[string]string)

	system["apple"] = "ios"
	system["microsoft"] = "windows"

	return system

}
