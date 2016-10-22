package main

import (
	"io"
	"log"
	"os"
)

// opens the files entered in command line and prints their contents
// if the files exits

func printFiles(str []string) {
	for _, val := range str {
		file, err := os.Open(val)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer file.Close()
		// copy the file to standard out
		// which prints it to the console
		io.Copy(os.Stdout, file)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("No filenames entered")
	}

	files := os.Args[1:]
	printFiles(files)
}
