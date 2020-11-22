package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	h := hash("hello world")
	fmt.Println(string(h))
	fmt.Println([]byte(string(h)))

	res := hex.EncodeToString(h)
	fmt.Println(res)

	deRes, err := hex.DecodeString(res)
	handle(err)

	fmt.Println(deRes)
}

func hash(data string) []byte {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(data)
	if err != nil {
		log.Fatal(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	return hash[:]
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
