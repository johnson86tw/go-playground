package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Header().Set("Johnson-Key", "this is from johnson")
	file, err := os.Open("index.json")
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(w, file)

	// fmt.Fprintln(w, "<p>Hello World</p>")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
