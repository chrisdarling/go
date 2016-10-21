package main

import (
	"fmt"
	"os"
	"strconv"
)

var conversion = map[string]float64{
	"mikm": 1.60934,
	"kmmi": 0.621371,
	/// add rest of distance types...
}

func convert(from string, to string) float64 {
	if len(from) < 2 {
		fmt.Println("Convert from metric not given.")
		os.Exit(1)
	}

	fromValue, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
	fromSuffix := from[len(from)-2:]
	key := fromSuffix + to

	if conversion[key] == 0 {
		fmt.Println("Metric does not exist or not supported")
		os.Exit(1)
	}

	return fromValue * conversion[key]

}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("missing arguments")
		os.Exit(1)
	}
	// Print the converted Result
	fmt.Println(convert(os.Args[1], os.Args[2]))

}
