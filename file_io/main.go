package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// open file
	file, err := os.Open("hello.txt")
	// how errors are checked in go
	if err != nil {
		// prints the message given and exits program
		log.Fatalln("File does not exist", err.Error())
	}
	// defers the operation until the program ends
	defer file.Close()
	// readAll returns bytes
	b, error := ioutil.ReadAll(file)
	if error != nil {
		log.Fatalln("operation failed!", err.Error())
	}

	// convert bytes to string
	str := string(b)

	// print contents of the file
	fmt.Println(str)

}
