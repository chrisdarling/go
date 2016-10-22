package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func mapWords(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	// splits into words
	scanner.Split(bufio.ScanWords)
	// every time Scan() is called it moves to the next ine of Text
	// and Text() gets the line of text
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		// for every instance of a word the value is incremented
		counts[word]++
	}

	return counts
}

func main() {

	src, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer src.Close()

	result := mapWords(src)

	//Prints instances of the looked up word in the file

	fmt.Println("Number of lorem's:", result["lorem"])

}
