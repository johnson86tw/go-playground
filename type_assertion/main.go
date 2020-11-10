package main

import "fmt"

func main() {
	args := []interface{}{1, "2", "3"}

	for _, arg := range args {
		if val, ok := arg.(int); ok {
			fmt.Println(val)
		} else {
			fmt.Println("false", val)
		}
	}
}
