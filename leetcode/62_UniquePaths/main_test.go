package main

import (
	"fmt"
	"testing"
)

var n int = 23

func Test_1(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("【input】:%v       【output】:%v\n", n, fibonacci(n))
	}

}

func Test_2(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("【input】:%v       【output】:%v\n", n, fibonacci2(n))
	}
}

func Test_factorial2(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("【input】:%v       【output】:%v\n", n, factorial2(n))
	}
}

func Test_factorial(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("【input】:%v       【output】:%v\n", n, factorial(n))
	}
}
