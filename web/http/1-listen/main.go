package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("form.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(w, req.Form)
}
