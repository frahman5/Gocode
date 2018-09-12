package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

var tpl *template.Template

type stockData struct {
	Date                                        time.Time
	OpenVal, High, Low, Close, Volume, AdjClose float64
}

func init() {
	var fM = map[string]interface{}{
		"df": dateFormat,
	}
	tpl = template.Must(template.New("").Funcs(fM).ParseFiles("tpl.gohtml"))
}

//dF, short for "date Format" formats the time for user display
func dateFormat (t time.Time) string {
	return t.Format("Jan 2 2006")
}

func main() {
	// Open the CSV file
	nf, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	// Read the csv into a [][]string
	reader := csv.NewReader(nf)
	reader.Read() // Discard the header line
	dataPasser, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dataPasser[0])

	// Pass all that data into a []stockData
	var data = make([]stockData, len(dataPasser))
	for k, v := range dataPasser {

		// create the time value
		t, err := time.Parse("2006-01-02", v[0])
		if err != nil {
			log.Fatalln(err)
		}

		// create the float values
		var floatPasser = make([]float64, 6)
		for l, m := range v[1:] {
			floatPasser[l], err = strconv.ParseFloat(m, 64)
			if err != nil {
				log.Fatalln(err)
			}
		}

		// add a value of type stockData to the data
		data[k] = stockData{
			t,
			floatPasser[0],
			floatPasser[1],
			floatPasser[2],
			floatPasser[3],
			floatPasser[4],
			floatPasser[5],
		}
	}

	// Create the Output File
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Execute the template
	err = tpl.ExecuteTemplate(f, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
