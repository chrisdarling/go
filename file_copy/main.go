package main

import (
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalln("missing arguments")
	}
	from := os.Args[1]
	to := os.Args[2]

	fromFile, err := os.Open(from)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer fromFile.Close()

	toFile, err := os.Create(to)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer toFile.Close()

	_, err = io.Copy(toFile, fromFile)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
