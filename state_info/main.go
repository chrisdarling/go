package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// struct corresponds to single row in csv file
type state struct {
	id                 int
	name               string
	abbreviation       string
	census_region_name string
}

func parseState(columns map[string]int, record []string) (*state, error) {
	col := columns["id"]
	id, err := strconv.Atoi(record[col])
	if err != nil {
		return nil, err
	}
	name := record[columns["name"]]
	abbreviation := record[columns["abbreviation"]]
	census_region_name := record[columns["census_region_name"]]

	return &state{
		id:                 id,
		name:               name,
		abbreviation:       abbreviation,
		census_region_name: census_region_name,
	}, nil

}

func main() {
	file, err := os.Open("../state_table.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	columns := make(map[string]int)
	stateLookup := make(map[string]state)
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err.Error())
		}
		// store columns in first row
		if rowCount == 0 {
			for i, col := range record {
				columns[col] = i
			}
		} else {
			state, err := parseState(columns, record)
			if err != nil {
				log.Fatalln(err.Error())
			}
			stateLookup[state.abbreviation] = *state
		}
	}

	if len(os.Args) < 2 {
		log.Fatalln("Enter state code as arguement")
	}

	// if ok is true state found
	// ok false state not found
	stateInfo, ok := stateLookup[os.Args[1]]
	if !ok {
		log.Fatalln("Invalid state code")
	}

	fmt.Println(stateInfo)

}
