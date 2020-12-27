package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is new Era.")
}

func f(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h3>台北流浪指南</h3>")
}

func main() {
	var h hotdog
	mux := http.NewServeMux()
	mux.Handle("/dog", h)

	mux.HandleFunc("/cat/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Meow")
	})

	// 第三種寫法
	mux.Handle("/taipei", http.HandlerFunc(f))

	http.ListenAndServe(":8080", mux)
}
