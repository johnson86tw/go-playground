package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 5))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	tmp := myPow(x, n/2)

	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}
