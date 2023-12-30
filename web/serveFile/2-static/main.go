package main

import (
	"fmt"
	"net/http"
)

func main() {
	// route must be "/" for root folder
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// 只有 "/" router 才會請求favicon.ico，"/dog"不會
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
}
