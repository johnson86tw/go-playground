package main

import (
	"fmt"
	"math/rand"
	"time"
)

type numList []int

func main() {
	numbers := []int{}

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for index, n := range numbers {
		if n%2 == 0 {
			fmt.Println(index, "is even")
		} else {
			fmt.Println(index, "is odd")
		}

	}

	foo()

}

func foo() {
	maxLength := 10
	n := numList{}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < maxLength; i++ {
		n = append(n, r.Intn(maxLength))
	}

	for _, v := range n {
		if v%2 == 0 {
			fmt.Printf("%v is even\n", v)
		} else {
			fmt.Printf("%v is odd\n", v)
		}
	}
}
