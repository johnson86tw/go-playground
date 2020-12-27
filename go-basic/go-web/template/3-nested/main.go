package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*"))
}

func main() {
	data := struct {
		Fake  int
		Truth int
	}{
		42,
		24,
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(res, "index.gohtml", data)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)

}
