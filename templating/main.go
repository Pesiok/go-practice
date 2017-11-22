package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       firstThree,
	"fdateMDY": monthDayYear,
	"fdbl":     double,
	"fsq":      square,
	"fsqrt":    squareRoot,
}

func init() {
	// parses and takes care of errors
	// tpl = template.Must(template.ParseGlob("templates/*"))

	// parses and takes care of errors & func map
	// funcs should be there before parse
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	parseHTML()
}

// template html funcs
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}
func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}
func double(x int) int {
	return x + x
}
func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}

// template methods
// then in template {{ .GetRandomValue }}
type person struct {
	Name string
	Age  int
}

func (p person) GetRandomValue() int {
	// #pdk
	return 4
}

func parseHTML() {

	// only one piece of data could be passed
	// data := time.Now()
	// data := []int{4, 5, 6, 7, 7, 9}
	data := 3
	err := tpl.ExecuteTemplate(os.Stdout, "go.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

	// err = tpl.ExecuteTemplate(os.Stdout, "mayn.gohtml", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
