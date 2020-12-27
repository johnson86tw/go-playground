package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog struct{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	mybook := req.FormValue("book")
	fmt.Println(mybook)

	data := struct {
		Method        string
		URL           *url.URL
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatalln(err)
	}
}
