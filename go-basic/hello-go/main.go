package main

import (
	"fmt"
	"math/big"
)

func main() {
	// ls := [][]string{{"hello", "hey", "cool"}, {"apple", "banana"}}
	// fmt.Println(ls

	// bs1 := []byte("hello")
	// bs2 := []byte("world")
	// info := bytes.Join([][]byte{bs1, bs2}, []byte{})
	// hash := sha256.Sum256(info)
	//
	// fmt.Println(hash)
	//
	// a := []string{"a", "b", "c"}
	// fmt.Println(a)

	newBigInt := big.NewInt(int64(123))
	fmt.Println(newBigInt)

}
