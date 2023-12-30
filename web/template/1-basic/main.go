package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	nf, err := os.Create("about.html")
	if err != nil {
		log.Fatalln(err)
	}

	defer nf.Close()
	defer fmt.Println("close the new file.")

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

	sli := []string{"a", "b", "c", "d"}
	m := map[string]string{
		"apple":     "ios",
		"microsoft": "windows",
	}
	data := struct {
		Num  int
		Name string
		Bar  []string
		M    map[string]string
	}{
		42,
		"Johnson",
		sli,
		m,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.gg", data)
	if err != nil {
		log.Fatalln(err)
	}

}
