package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type record struct {
	Date time.Time
	Open float64
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	data := parseCSV("table.csv")

	tpl := template.Must(template.ParseFiles("index.gohtml"))
	err := tpl.Execute(res, data)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseCSV(filename string) []record {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	records := make([]record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, record{
			Date: date,
			Open: open,
		})
	}

	return records
}
