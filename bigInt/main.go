package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

func main() {

}

func hashAndBigInt() {
	hash := sha256.Sum256([]byte("qf Wqrreqwrorld\n"))
	fmt.Println(len(hash))

	var intHash big.Int
	intHash.SetBytes(hash[:])

	target := big.NewInt(1)
	target.Lsh(target, uint(256))

	// fmt.Println("0b", target.Text(2))
	fmt.Println(target.Text(10))
	fmt.Println(target.Text(16), "base16")
	fmt.Println(intHash.Text(16))

}

func endianness() {
	v := uint32(500)
	littlebuf := make([]byte, 4)
	bigbuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(littlebuf, v)
	binary.BigEndian.PutUint32(bigbuf, v)
	fmt.Println(littlebuf)
	fmt.Println(bigbuf)
	// 似乎不適合直接印出來，要使用 Write，因為 buff 仍是 stream
}

func howToUseBytesJoin() {
	h := []byte("Hello")
	w := []byte("World")

	s := bytes.Join([][]byte{h, w}, []byte(" "))
	fmt.Println(string(s))
}

func convertInt64ToByteSlice(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func aboutBigInt() {
	newBigInt := big.NewInt(int64(-10))
	// convert big int to unsigned integer 64 byte
	fmt.Println(newBigInt.Uint64())

	// convert big int to string in hex
	fmt.Println(newBigInt.Text(16))

	fmt.Printf("%d ~ %d\n", math.MinInt8, math.MaxInt8)

	//use SetString convert string in base into big int
	i, _ := new(big.Int).SetString("A", 16)
	fmt.Println(i)
}
