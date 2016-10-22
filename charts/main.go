package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type record struct {
	Date string
	Open float64
}

func makeRecord(row []string) record {
	open, _ := strconv.ParseFloat(row[1], 64)
	return record{
		Date: row[0],
		Open: open,
	}
}

func main() {
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	rdr := csv.NewReader(file)

	rows, err := rdr.ReadAll()

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(`<!Doctype html>
    <html>
    <head></head>
    <body>
      <table>
        <thead>
          <tr>
            <th>Date</th>
            <th>Open</th>
          </tr>
        </thead>
        <tbody>
      `)

	for i, row := range rows {
		if i != 0 {
			record := makeRecord(row)
			fmt.Println(`
          <tr>
            <td>` + record.Date + `</td>
            <td>` + fmt.Sprintf("%0.2f", record.Open) + `</td>
          </tr>`)
		}

	}

	fmt.Println(`</tbody>

  </table>`)
}
