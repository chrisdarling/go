package main

import (
	"log"
	"os"
)

func main() {
	// create file
	file, err := os.Create("hello.txt")
	// how errors are checked in go
	if err != nil {
		// prints the message given and exits program
		log.Fatalln("File does not exist", err.Error())
	}
	// defers the operation until the program ends
	defer file.Close()

	// write to the file
	_, err = file.Write([]byte("Write some text here"))

	if err != nil {
		log.Fatalln("error when writing to the file", err.Error())
	}

}
