package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

type user struct {
	Name  string
	Admin bool
	Age   int
}

type data struct {
	Name  string
	Thing string
	T     time.Time
	Users []user
}

var tpl *template.Template
var fm = template.FuncMap{
	"uc":    strings.ToUpper,
	"ft":    firstThree,
	"tf":    timef,
	"fsq":   square,
	"fsqrt": sqRoot,
}
var u1 = user{"John", false, 22}
var u2 = user{"Grace", true, 20}
var u3 = user{"Patric", false, 21}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	users := []user{u1, u2, u3}
	d := data{
		"john",
		"pencil",
		time.Now(),
		users,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", d)
	if err != nil {
		log.Fatalln(err)
	}
}

func (d data) Number() int {
	return 42
}

func (d data) AddLastName(s string) string {
	return d.Name + s
}

func (u user) DoubleAge() int {
	return u.Age * u.Age
}

func firstThree(s string) string {
	if len(s) > 3 {
		return s[:3]
	}
	return s
}

func timef(t time.Time) string {
	return t.Format(`Jan 02, 06`)
}

func square(i int) float64 {
	return float64(i * i)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}
